package model

import (
	"encoding/json"
	"fmt"
)

var (
	// general error
	ErrNil            = &BusinessError{Code: 0, Message: "OK"}
	ErrValidation     = &BusinessError{Code: 1, Message: "Invalid parameter"}
	ErrInternalServer = &BusinessError{Code: 2, Message: "Internal server error"}
	ErrAuth           = &BusinessError{Code: 3, Message: "Authentication failed"}

	// specific error
	ErrAppNotFound        = &BusinessError{Code: 10000, Message: "APP not found"}
	ErrResourceNotFound   = &BusinessError{Code: 10001, Message: "Resource not found"}
	ErrInsufficentBalance = &BusinessError{Code: 10002, Message: "Insufficient balance"}
	ErrAccountFrozen      = &BusinessError{Code: 10003, Message: "Account fronzen"}
	ErrInvalidAppOperator = &BusinessError{Code: 10004, Message: "Invalid APP operator"}
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
	return fmt.Sprintf("%s: %v", err.Message, err.Data)
}

func (err *BusinessError) IsNil() bool {
	return err == nil || err.Code == ErrNil.Code
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

func IsBusinessError(err error) (*BusinessError, bool) {
	be, ok := err.(*BusinessError)
	return be, ok
}
