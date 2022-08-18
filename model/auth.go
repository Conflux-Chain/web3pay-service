package model

const (
	AuthHeaderBillingKey  = "Billing-Key"
	AuthHeaderCustomerKey = "Customer-Key"
)

type AuthKeyObject struct {
	Msg string `json:"msg"` // signed message
	Sig string `json:"sig"` // signature
}
