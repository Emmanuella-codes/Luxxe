package dtos

type CreateProductDTO struct {
	UserID      string  `json:"userID" validate:"required"`
	Name        string 	`json:"name" validate:"required"`
	Description string 	`json:"description" validate:"required,min=10"`
	Category 		string 	`json:"category" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Quantity 		int			`json:"quantity" validate:"required"`
}
