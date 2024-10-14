package dtos

type SendOTPDTO struct {
	Email string `json:"email" validate:"required,email"`
}
