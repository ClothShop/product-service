package dto

import "mime/multipart"

type ProductCreate struct {
	Name        string                  `form:"product_name" binding:"required"`
	Description string                  `form:"description" binding:"required"`
	Price       float64                 `form:"price" binding:"required,gt=0"`
	Images      []*multipart.FileHeader `form:"images" binding:"required"`
}
