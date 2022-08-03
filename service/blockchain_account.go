package service

import (
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// DeleteAccountStatus deletes APP coin status (balance, frozen status etc.,) of specific address.
func (svc *BlockchainService) DeleteAccountStatus(appCoin, address common.Address) (*model.AppCoinAccount, error) {
	lockKey := util.MutexKey(appCoin.String() + address.String())
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

	account, _, err := svc.memStore.DeleteAccount(appCoin, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get App coin account")
	}

	return account, nil
}

// UpdateAccountStatus updates APP coin status (balance, frozen status etc.,) of specific address.
func (svc *BlockchainService) UpdateAccountStatus(
	appCoin, address common.Address, balance, frozen, block *int64) (*model.AppCoinAccount, error) {
	lockKey := util.MutexKey(appCoin.String() + address.String())
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

	account, ok, err := svc.memStore.GetAccount(appCoin, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get App coin account")
	}

	if !ok {
		return nil, errors.New("APP coin account not exited")
	}

	logger := logrus.WithFields(logrus.Fields{
		"appCoin": appCoin, "address": address,
		"frozen": frozen, "block": block, "balance": balance,
	})

	if frozen != nil {
		account.Frozen = *frozen
	}

	if balance != nil {
		account.Balance = *balance
	}

	if block != nil {
		account.ConfirmedBlock = *block

		copy := *account
		if err := svc.memStore.SaveAccount(&copy); err != nil {
			return nil, errors.WithMessage(err, "failed to save APP coin account")
		}
	}

	logger.WithField("account", account).Debug("App coin account status updated")
	return account, nil
}

// GetAccountStatus gets APP coin status (balance, frozen status etc.,) of specific address.
func (svc *BlockchainService) GetAccountStatus(appCoin, address common.Address) (*model.AppCoinAccount, error) {
	logger := logrus.WithFields(logrus.Fields{
		"appCoin": appCoin,
		"address": address,
	})

	account, ok, err := svc.memStore.GetAccount(appCoin, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get App coin account")
	}

	if ok {
		logger.WithField("account", account).Debug("App coin account existed already")
		return account, nil
	}

	lockKey := util.MutexKey(appCoin.String() + address.String())
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

	// double checking
	account, ok, err = svc.memStore.GetAccount(appCoin, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get App coin account")
	}

	if ok {
		logger.WithField("account", account).Debug("App coin account existed already")
		return account, nil
	}

	balance, frozen, err := svc.provider.GetAppCoinBalanceAndFrozenStatus(nil, appCoin, address)
	if err != nil {
		logger.WithError(err).Info("Failed to get APP coin balance/frozen status")
		return nil, errors.WithMessage(err, "failed to get APP coin balance/frozen status")
	}

	coin, addr := appCoin.Hex(), address.Hex()
	bill, existed, err := svc.sqliteStore.GetBill(coin, addr)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP coin bill")
	}

	account = model.NewAppCoinAccount(coin, addr, frozen, balance)
	if existed {
		account.IncreaseFee(bill.Fee)
	}

	if err := svc.memStore.SaveAccount(account); err != nil {
		return nil, errors.WithMessage(err, "failed to save APP coin account")
	}

	svc.addStatusConfirmTask(appCoin, address)

	logger.WithField("appCoinAccount", account).Debug("Created APP coin account for address")
	return account, nil
}

// DeductAccountBalance deducts APP coin balance of specific account.
func (svc *BlockchainService) DeductAccountBalance(appCoin, address common.Address, amount int64) error {
	lockKey := util.MutexKey(appCoin.String() + address.String())
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

	account, ok, err := svc.memStore.GetAccount(appCoin, address)
	if err != nil {
		return errors.WithMessage(err, "failed to get APP coin account")
	}

	if !ok {
		return errors.New("APP coin account not existed")
	}

	if account.TotalBalance() < int64(amount) {
		return model.ErrInsufficentBalance
	}

	account.IncreaseFee(amount)
	return nil
}
