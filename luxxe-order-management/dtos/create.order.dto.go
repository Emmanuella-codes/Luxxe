package dtos

type CreateOrderDTO struct {
	UserID          string `json:"userID" validate:"required"`
	ShippingAddress string `json:"shippingAddress"`
	PhoneNumber     int64  `json:"phoneNumber"`
}
