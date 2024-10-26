package dtos

type GetCartDTO struct {
	UserID    string `query:"userID" validate:"required"`
}
