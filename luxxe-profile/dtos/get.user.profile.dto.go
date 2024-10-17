package dtos

type GetUserProfileDTO struct {
	UserID string `validate:"required,min=24,max=24"`
}
