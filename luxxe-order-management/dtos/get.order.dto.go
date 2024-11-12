package dtos

type GetOrderDTO struct {
	UserID  string `json:"userID" validate:"required"`
}
