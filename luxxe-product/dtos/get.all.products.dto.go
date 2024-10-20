package dtos

type GetProductDTO struct {
	Page 	 int 		`query:"page" validate:"required"`
	// UserID string `query:"ID" validate:"required,min=24,max=24"`
}