package dtos

type VerifyEmailDTO struct {
	UserID string `json:"userID" validate:"required,min=24,max=24"`
}
