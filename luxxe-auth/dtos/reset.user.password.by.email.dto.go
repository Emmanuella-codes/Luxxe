package dtos

type ResetUserPasswordByEmailDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Otp      string `json:"otp" validate:"required,min=4,max=4"`
	Password string `json:"password" validate:"required,min=8"`
}
