package dtos

type AddToCartDTO struct {
	UserID    string  `json:"userID" validate:"required"`
	ProductID string  `json:"productID" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price     float64 `json:"price"`
}
