package dtos

import entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"

type GetProductByCategoryDTO struct {
	Page     int            						`query:"page" validate:"required"`
	Category entities.ProductCategories `query:"category" validate:"required"`
}
