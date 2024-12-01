package dtos

type CreateOrderDTO struct {
	UserID          string `json:"userID" validate:"required"`
	CartID          string `json:"cartID" validate:"required"`
	ShippingAddress string `json:"shippingAddress"`
	PhoneNumber     string  `json:"phoneNumber"`
}
