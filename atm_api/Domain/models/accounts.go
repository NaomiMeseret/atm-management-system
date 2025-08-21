package models
 import(
	"Time"
 )
type Account struct{
	ID string
	UserID string
	BalanceCents  int64
	Currency string
	CreatedAt time.Time
	UpdatedAt time.Time
}