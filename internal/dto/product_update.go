package dto

type ProductUpdate struct {
	ID          uint     `json:"id" validate:"required"`
	Name        *string  `json:"name,omitempty" validate:"omitempty,required"`
	Description *string  `json:"description,omitempty" validate:"omitempty,required"`
	Price       *float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
}
