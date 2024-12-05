package dtos

type UpdateOrderDTO struct {
	UserID          string `json:"userID" validate:"required"`
	CartID          string `json:"cartID" validate:"required"`
	OrderID         string `json:"orderID" validate:"required"`
	ShippingAddress string `json:"shippingAddress"`
	PhoneNumber     string `json:"phoneNumber"`
}
