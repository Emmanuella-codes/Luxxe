package dtos

type UpdateCartItemDTO struct {
	UserID    string `json:"userID" validate:"required"`
	ProductID string `json:"productID" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}
