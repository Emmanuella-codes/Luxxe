package dtos

type ResetUserPasswordByUserIDDTO struct {
	UserID   string `json:"userID"  validate:"required,min=24,max=24"`
	Otp      string `json:"otp" validate:"required,min=4,max=4"`
	Password string `json:"password" validate:"required,min=8"`
}
