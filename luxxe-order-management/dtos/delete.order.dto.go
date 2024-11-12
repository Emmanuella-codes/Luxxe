package dtos

type DeleteOrderDTO struct {
	UserID   string `json:"userID" validate:"required"`
}
