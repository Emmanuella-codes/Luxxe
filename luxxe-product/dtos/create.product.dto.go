package dtos

import entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"

type CreateProductDTO struct {
	UserID      	string     									`json:"userID" validate:"required"`
	Name        	string     									`json:"name" validate:"required"`
	Description 	string     									`json:"description" validate:"required,min=10"`
	Category    	entities.ProductCategories 	`json:"category" validate:"required,oneof=hair skincare fragrance wellness makeup accessories personalcare bath&body jewelry homespa"`
	Price       	float64    									`json:"price" validate:"required,gt=0"`
	ProductImage 	string											`json:"productImage" validate:"required"`
	ProductInfo 	string											`json:"productInfo" validate:"required"`
	Quantity    	int        									`json:"quantity" validate:"required"`
}
