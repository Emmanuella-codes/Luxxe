package dtos

type ClearCartDTO struct {
	UserID    string `json:"userID" validate:"required"`
}
