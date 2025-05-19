package product

import "mime/multipart"

type Create struct {
	Name           string                  `json:"name" validate:"required"`
	Description    string                  `json:"description" validate:"required"`
	Price          float64                 `json:"price" validate:"required,gt=0"`
	CompareAtPrice float64                 `json:"compare_at_price" validate:"omitempty,required"`
	Images         []*multipart.FileHeader `form:"image_files" binding:"required"`
	CategoryID     uint                    `json:"category_id" validate:"required"`
}
