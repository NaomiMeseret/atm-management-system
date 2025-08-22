package dto

import "time"

type TransactionResponse struct{
	Type string `json:"type"`
	Amount float64 `json:"amount"`
	Balance float64 `json:"balance"`
	Counterparty *string `json:"counterparty,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}