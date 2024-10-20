package dtos

type DeleteProductDTO struct {
	ProductID 	string `json:"productID" validate:"required"`
}
