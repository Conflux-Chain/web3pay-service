package model

var (
	// general error
	ErrNil            = &BusinessError{Code: 0, Message: "OK"}
	ErrValidation     = &BusinessError{Code: 1, Message: "Invalid parameter"}
	ErrInternalServer = &BusinessError{Code: 2, Message: "Internal server error"}
	ErrAuth           = &BusinessError{Code: 3, Message: "Authentication failed"}

	// specific error
	ErrAppCoinNotFound            = &BusinessError{Code: 10000, Message: "APP coin not found"}
	ErrAppCoinResourceNotFound    = &BusinessError{Code: 10001, Message: "APP coin resource not found"}
	ErrInsufficentBalance         = &BusinessError{Code: 10002, Message: "Insufficient balance"}
	ErrAccountAddrFrozen          = &BusinessError{Code: 10003, Message: "Account address fronzen"}
	ErrNotAnValidAppCoinOwner     = &BusinessError{Code: 10004, Message: "Not a valid APP coin contract owner"}
	ErrAppCoinAddrBalanceNotFound = &BusinessError{Code: 10005, Message: "address balance of APP coin not found"}
)

type BusinessError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (be *BusinessError) WithData(data interface{}) *BusinessError {
	return &BusinessError{
		Code: be.Code, Message: be.Message, Data: data,
	}
}

func (err *BusinessError) Error() string {
	return err.Message
}
