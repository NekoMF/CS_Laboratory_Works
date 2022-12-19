package dto

type CreateUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Otp      string `json:"otp" binding:"required"`
}
