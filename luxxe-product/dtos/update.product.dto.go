package dtos

import entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"

type UpdateProductDTO struct {
	ProductID   	string     									`json:"productID" validate:"required"`
	Name        	string     									`json:"name"`
	Description 	string     									`json:"description" validate:"min=10"`
	Category    	entities.ProductCategories  `json:"category" validate:"oneof=hair skincare fragrance wellness makeup accessories personalcare bath&body jewelry homespa"`
	Price       	float64    									`json:"price" validate:"gt=0"`
	ProductImage 	string											`json:"productImage"`
	ProductInfo 	string											`json:"productInfo"`
	Quantity    	int        									`json:"quantity"`
}
