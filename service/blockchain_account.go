package service

import (
	"math/big"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

// DeleteAccountStatus deletes APP coin status (balance, frozen status etc.,) of specific address.
func (svc *BlockchainService) DeleteAccountStatus(appCoin, address common.Address) (*model.AppCoinAccount, error) {
	lockKey := util.MutexKey(appCoin.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	account, _, err := svc.memStore.DeleteAccount(appCoin, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get App coin account")
	}

	return account, nil
}

// UpdateAccountStatus updates APP coin status (balance, frozen status etc.,) of specific address.
func (svc *BlockchainService) UpdateAccountStatus(
	appCoin, address common.Address, balance *big.Int, frozen, block *int64) (*model.AppCoinAccount, error) {
	lockKey := util.MutexKey(appCoin.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	logger := logrus.WithFields(logrus.Fields{
		"appCoin": appCoin,
		"address": address,
		"block":   block,
		"frozen":  frozen,
	})

	account, ok, err := svc.memStore.GetAccount(appCoin, address)
	if err != nil {
		logger.WithError(err).Info("Failed to get APP coin account for status update")
		return nil, errors.WithMessage(err, "failed to get App coin account")
	}

	if !ok {
		return nil, nil
	}

	if frozen != nil {
		account.Frozen = *frozen
	}

	if balance != nil {
		logger = logger.WithField("newBalance", balance.String())
		account.Balance = decimal.NewFromBigInt(balance, 0)
	}

	if block != nil {
		account.ConfirmedBlock = *block

		// re-index
		copy := *account
		if err := svc.memStore.SaveAccount(&copy); err != nil {
			logger.WithError(err).Info("Failed to save APP coin account for status update")
			return nil, errors.WithMessage(err, "failed to save APP coin account")
		}
	}

	logger.WithField("account", account).Debug("App coin account status updated")
	return account, nil
}

// GetOrFetchAccountStatus gets APP coin status (balance, frozen status etc.,) of specific address.
func (svc *BlockchainService) GetOrFetchAccountStatus(appCoin, address common.Address) (*model.AppCoinAccount, error) {
	account, ok, err := svc.memStore.GetAccount(appCoin, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP coin account")
	}

	if ok {
		return account, nil
	}

	lockKey := util.MutexKey(appCoin.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	// double checking
	account, ok, err = svc.memStore.GetAccount(appCoin, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP coin account")
	}

	if ok {
		return account, nil
	}

	logrus.WithFields(logrus.Fields{
		"coin": appCoin, "address": address,
	}).Debug("Fetch APP coin account status due not existed yet")

	// fetch balance and frozen status
	balance, frozen, err := svc.provider.GetAppCoinBalanceAndFrozenStatus(nil, appCoin, address)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to fetch APP coin balance/frozen status")
	}

	// load bill fee
	coin, addr := appCoin.String(), address.String()
	fee, err := svc.sqliteStore.GetBillFee(coin, addr)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to load bill fee")
	}

	// save account status
	account = model.NewAppCoinAccount(coin, addr, frozen, balance)
	account.IncreaseFee(fee)

	if err := svc.memStore.SaveAccount(account); err != nil {
		return nil, errors.WithMessage(err, "failed to save APP coin account")
	}

	// send to status confirmation queue
	svc.addStatusConfirmTask(appCoin, address)

	logrus.WithField("appCoinAccount", account).Debug("APP coin account fetched and created")
	return account, nil
}

// DeductAccountBalance deducts APP coin balance of specific account.
func (svc *BlockchainService) DeductAccountBalance(appCoin, address common.Address, amount *big.Int) (bool, error) {
	lockKey := util.MutexKey(appCoin.String() + address.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	account, ok, err := svc.memStore.GetAccount(appCoin, address)
	if err != nil {
		return false, errors.WithMessage(err, "failed to get APP coin account")
	}

	if !ok {
		return false, nil
	}

	if account.TotalBalance().Cmp(amount) <= 0 {
		return false, model.ErrInsufficentBalance
	}

	account.IncreaseFee(amount)
	return true, nil
}
