package models
import(
	"Time"
)

type Transaction struct{
	ID string
	AccountID string
	Type string
	AmountCents int64
	BalanceAfterCents  int64
	Counterparty *string
	CreatedAt time.Time
}