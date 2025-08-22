package dto

type AmountRequest struct{
	Amount float64 `json:"amount"`
}

type BalanceResponse struct{
	AccountNumber  string `json:"account_number"`
	Balance float64 `json:"balance"`
}
