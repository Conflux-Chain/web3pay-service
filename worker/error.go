package worker

import "regexp"

var (
	errPatternTxnAlreadyExists      = regexp.MustCompile(`(.*)tx already exist(.*)`)
	errPatternTxnNonceAlreadyExists = regexp.MustCompile(`(.*)with same nonce already inserted(.*)`)
	errPatternTxnNonceTooStale      = regexp.MustCompile(`(.*)discarded due to a too stale nonce(.*)`)
	errPatternTxnNonceTooFuture     = regexp.MustCompile(`(.*)discarded due to in too distant future(.*)`)
	errPatternTxnPoolIsFull         = regexp.MustCompile(`(.*)Transaction Pool is full(.*)`)
	errPatternTxnPoolIsFull2        = regexp.MustCompile(`(.*)txpool is full(.*)`)
	errPatternTxnGasTooSmall        = regexp.MustCompile(`(.*)NotEnoughBaseGas(.*)`)
	errPatternTxnGasTooLarge        = regexp.MustCompile(`(.*)transaction gas \d+ exceeds the maximum value(.*)`)
	errPatternTxnGasPriceIsZero     = regexp.MustCompile(`(.*)ZeroGasPrice(.*)`)
)

func isTxnAlreadyExistError(err error) bool {
	return errPatternTxnAlreadyExists.MatchString(err.Error())
}

func isTxnPollFullError(err error) bool {
	return errPatternTxnPoolIsFull.MatchString(err.Error()) || errPatternTxnPoolIsFull2.MatchString(err.Error())
}

func isTxnNonceAlreadyExistsError(err error) bool {
	return errPatternTxnNonceAlreadyExists.MatchString(err.Error())
}

func isTxnNonceTooStaleError(err error) bool {
	return errPatternTxnNonceTooStale.MatchString(err.Error())
}

func isTxnNonceTooFutureError(err error) bool {
	return errPatternTxnNonceTooFuture.MatchString(err.Error())
}

func isTxnGasTooSmallError(err error) bool {
	return errPatternTxnGasTooSmall.MatchString(err.Error())
}

func isTxnGasTooLargeError(err error) bool {
	return errPatternTxnGasTooLarge.MatchString(err.Error())
}

func isTxnGasPriceIsZeroError(err error) bool {
	return errPatternTxnGasPriceIsZero.MatchString(err.Error())
}
