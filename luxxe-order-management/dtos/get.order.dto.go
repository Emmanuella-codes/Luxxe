package dtos

type GetOrderDTO struct {
	UserID  string `query:"userID" validate:"required"`
}
