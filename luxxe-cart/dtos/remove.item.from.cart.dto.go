package dtos

type RemoveItemFromCartDTO struct {
	UserID    string `json:"userID" validate:"required"`
	ProductID string `json:"productID" validate:"required"`
}
