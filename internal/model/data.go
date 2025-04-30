package model

type Account struct {
	AccountName string  `json:"account_name"`
	Balance     float32 `json:"account_balance"`
}
