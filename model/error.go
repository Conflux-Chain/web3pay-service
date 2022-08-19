package model

import "encoding/json"

var (
	// general error
	ErrNil            = &BusinessError{Code: 0, Message: "OK"}
	ErrValidation     = &BusinessError{Code: 1, Message: "Invalid parameter"}
	ErrInternalServer = &BusinessError{Code: 2, Message: "Internal server error"}
	ErrAuth           = &BusinessError{Code: 3, Message: "Authentication failed"}

	// specific error
	ErrAppCoinNotFound         = &BusinessError{Code: 10000, Message: "APP coin not found"}
	ErrAppCoinResourceNotFound = &BusinessError{Code: 10001, Message: "APP coin resource not found"}
	ErrInsufficentBalance      = &BusinessError{Code: 10002, Message: "Insufficient balance"}
	ErrAccountAddrFrozen       = &BusinessError{Code: 10003, Message: "Account address fronzen"}
	ErrNotAnValidAppCoinOwner  = &BusinessError{Code: 10004, Message: "Not a valid APP coin contract owner"}
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

// GetObject converts the business error data to an arbitrary type.
//
// The function works as you would expect it from json.Unmarshal()
func (be *BusinessError) GetObject(toType interface{}) error {
	js, err := json.Marshal(be.Data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, toType)
	if err != nil {
		return err
	}

	return nil
}
