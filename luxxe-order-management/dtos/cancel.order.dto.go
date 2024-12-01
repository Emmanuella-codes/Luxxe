package dtos

type CancelOrderDTO struct {
	UserID   string `json:"userID" validate:"required"`
}
