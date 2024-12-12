package dtos

import entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"

type CreateOrderDTO struct {
	UserID          string 							 `json:"userID" validate:"required"`
	CartID          string 							 `json:"cartID" validate:"required"`
	Email           string               `json:"email"`
	ShippingAddress string 							 `json:"shippingAddress"`
	PhoneNumber     string 							 `json:"phoneNumber"`
	OrderStatus     entities.OrderStatus `json:"orderStatus"`
}
