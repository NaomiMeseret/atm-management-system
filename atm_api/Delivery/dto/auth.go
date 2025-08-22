package dto

type SignupRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct{
	Message string `json:"message"`
	AccountNumber string `json:"account_number"`
}