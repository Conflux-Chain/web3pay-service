package api

var (
	// general error
	errNil            = &businessError{Code: 0, Message: "ok"}
	errValidation     = &businessError{Code: 1, Message: "invalid parameter"}
	errInternalServer = &businessError{Code: 2, Message: "internal server error"}
)

type businessError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (be *businessError) withData(data interface{}) *businessError {
	return &businessError{
		Code: be.Code, Message: be.Message, Data: data,
	}
}

func (err *businessError) Error() string {
	return err.Message
}
