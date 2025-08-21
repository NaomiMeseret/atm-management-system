package models

import "Time"

type User struct{
	ID string
	Username string
	Password string
	AccountNumber string
	CreatedAt time.Time
	
}