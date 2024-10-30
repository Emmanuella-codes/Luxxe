package dtos

type GetCartDTO struct {
	UserID string `json:"userID" validate:"required"`
	Page 	 int 		`query:"page" validate:"required"`
}
