package dtos

import entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"

type CartItemsDTO struct {
	ProductID string  `json:"productID" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
}

type CartDTO struct {
	CartID    string  				`json:"cartID" validate:"required"`
	CartItems []CartItemsDTO 	`json:"cartItems"`
}

type PlaceOrderDTO struct {
	CartDTO
	UserID        string 							    `json:"userID" validate:"required"`
	PaymentClient entities.PaymentClient  `json:"paymentClient" validate:"required,oneof=paystack"`
	OrderID       string                  `json:"orderID" validate:"required"`
}
