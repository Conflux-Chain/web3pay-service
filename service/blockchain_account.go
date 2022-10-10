package service

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

// DeleteAccountStatus deletes APP account of specific address.
func (svc *BlockchainService) DeleteAccountStatus(app, address common.Address) (*model.AppAccount, error) {
	lockKey := util.MutexKey(app.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	account, _, err := svc.memStore.DeleteAccount(app, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to delete App account")
	}

	return account, nil
}

// UpdateAccountStatus updates APP account status (balance, frozen status etc.,) of specific address.
func (svc *BlockchainService) UpdateAccountStatus(
	app, address common.Address, balance *big.Int, frozen, block *int64) (*model.AppAccount, error) {
	lockKey := util.MutexKey(app.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	account, ok, err := svc.memStore.GetAccount(app, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get App account")
	}

	if !ok {
		return nil, nil
	}

	if frozen != nil {
		account.Frozen = *frozen
	}

	if balance != nil {
		account.Balance = decimal.NewFromBigInt(balance, 0)
	}

	if block != nil {
		account.ConfirmedBlock = *block

		// re-index
		copy := *account
		if err := svc.memStore.SaveAccount(&copy); err != nil {
			return nil, errors.WithMessage(err, "failed to save APP account")
		}
	}

	return account, nil
}

// GetOrFetchAppAccount gets APP account status of specific address.
func (svc *BlockchainService) GetOrFetchAppAccountStatus(app, address common.Address) (*model.AppAccount, error) {
	account, ok, err := svc.memStore.GetAccount(app, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP account")
	}

	if ok {
		return account, nil
	}

	lockKey := util.MutexKey(app.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	// double checking
	account, ok, err = svc.memStore.GetAccount(app, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP account")
	}

	if ok {
		return account, nil
	}

	// fetch balance and frozen status
	balance, frozen, err := svc.provider.GetAppAccountBalanceAndFrozenStatus(nil, app, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to fetch APP account balance/frozen status")
	}

	// load bill fee
	fee, err := svc.sqliteStore.GetBillFee(app.String(), address.String())
	if err != nil {
		return nil, errors.WithMessage(err, "failed to load bill fee")
	}

	// save account status
	account = model.NewAppAccount(app.String(), address.String(), frozen, balance)
	account.IncreaseFee(fee)

	txn := svc.memStore.Txn(true)
	if err := svc.memStore.SaveAccountWithTxn(txn, account); err != nil {
		txn.Abort()
		return nil, errors.WithMessage(err, "failed to save APP account")
	}

	// send to status confirmation queue
	if svc.addStatusConfirmTask(app, address) {
		txn.Commit()
	} else {
		txn.Abort()
	}

	logrus.WithField("appAccount", account).Debug("APP account fetched and created")
	return account, nil
}

// WriteOffAccountFee writes off APP fee from specific account.
func (svc *BlockchainService) WriteOffAccountFee(app, address common.Address, amount *big.Int) (bool, error) {
	lockKey := util.MutexKey(app.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	account, ok, err := svc.memStore.GetAccount(app, address)
	if err != nil {
		return false, errors.WithMessage(err, "failed to get APP account")
	}

	if !ok {
		return false, nil
	}

	account.DecreaseFee(amount)
	account.DecreaseBalance(amount)

	return true, nil
}

// WithholdAccountFee withholds APP fee from specific account.
func (svc *BlockchainService) WithholdAccountFee(app, address common.Address, amount *big.Int) (bool, error) {
	lockKey := util.MutexKey(app.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	account, ok, err := svc.memStore.GetAccount(app, address)
	if err != nil {
		return false, errors.WithMessage(err, "failed to get APP account")
	}

	if !ok {
		return false, nil
	}

	if account.TotalBalance().Cmp(amount) < 0 {
		return false, model.ErrInsufficentBalance
	}

	account.IncreaseFee(amount)
	return true, nil
}

// IncreaseAccountBalance increases APP balance of specific account.
func (svc *BlockchainService) IncreaseAccountBalance(
	app, address common.Address, amount *big.Int, block int64) (bool, error) {

	return svc.changeAccountBalance(app, address, amount, block)
}

// DecreaseAccountBalance decreases APP balance of specific account.
func (svc *BlockchainService) DecreaseAccountBalance(
	app, address common.Address, amount *big.Int, block int64) (bool, error) {

	return svc.changeAccountBalance(app, address, amount, block, true)
}

func (svc *BlockchainService) changeAccountBalance(
	app, address common.Address, amount *big.Int, block int64, decrease ...bool) (bool, error) {

	lockKey := util.MutexKey(app.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	account, ok, err := svc.memStore.GetAccount(app, address)
	if err != nil {
		return false, errors.WithMessage(err, "failed to get APP account")
	}

	if !ok {
		return false, nil
	}

	if !account.IsConfirmed() { // not confirmed yet?
		return false, nil
	}

	if account.ConfirmedBlock >= block { // stale block?
		return false, nil
	}

	if len(decrease) > 0 && decrease[0] {
		account.DecreaseBalance(amount)
	} else {
		account.IncreaseBalance(amount)
	}

	return true, nil
}

// checkOperatorBalance periodically check operator balance.
func (bs *BlockchainService) checkOperatorBalance() {
	var config struct {
		OperatorBalanceCheckInterval  time.Duration `default:"30m"`
		OperatorBalanceCheckThreshold int64         `default:"15"`
	}
	viper.MustUnmarshalKey("blockchain", &config)

	bs.checkOperatorBalanceOnce(config.OperatorBalanceCheckThreshold)

	ticker := time.NewTicker(config.OperatorBalanceCheckInterval)
	defer ticker.Stop()

	for range ticker.C {
		err := bs.checkOperatorBalanceOnce(config.OperatorBalanceCheckThreshold)
		if err != nil {
			logrus.WithError(err).Error("Failed to check operator balance once")
			ticker.Reset(config.OperatorBalanceCheckInterval / 2)
		} else {
			ticker.Reset(config.OperatorBalanceCheckInterval)
		}
	}
}

func (bs *BlockchainService) checkOperatorBalanceOnce(threshold int64) error {
	blockNum := types.BlockNumberOrHashWithNumber(types.LatestBlockNumber)
	operatorAddr := bs.provider.OperatorAddress()

	balance, err := bs.provider.Balance(operatorAddr, &blockNum)
	if err != nil {
		logrus.WithField("operatorAddr", operatorAddr).WithError(err).Info("Failed to check operator CFX blance")
		return err
	}

	thresholdD := decimal.NewFromInt(threshold).Mul(decimal.New(1, 18))
	balanceD := decimal.NewFromBigInt(balance, 0)

	logger := logrus.WithFields(logrus.Fields{
		"operatorAddr": operatorAddr,
		"balance":      balanceD.Div(decimal.New(1, 18)).StringFixed(2),
		"threshold":    thresholdD.Div(decimal.New(1, 18)).StringFixed(2),
	})

	if balanceD.Cmp(thresholdD) <= 0 {
		logger.Warn("The balance for the operater is too low")
	} else {
		logger.Debug("The operator balance checked once")
	}

	return nil
}
