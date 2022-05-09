package model

type PasswordReset struct {
	Code     string `json:"code"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
