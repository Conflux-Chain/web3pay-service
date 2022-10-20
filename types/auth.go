package types

const (
	AuthHeaderBillingKey = "Billing-Key"
	AuthHeaderApiKey     = "Api-Key"
)

type BillingAuthKey struct {
	Msg string `json:"msg"` // signed message
	Sig string `json:"sig"` // signature
}

type ApiAuthKey struct {
	Sig       string // signature
	Timestamp int64  // timestamp
}
