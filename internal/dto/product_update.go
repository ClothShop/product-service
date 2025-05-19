package dto

type ProductUpdate struct {
	ID             uint     `json:"id" validate:"required"`
	Name           *string  `json:"name,omitempty" validate:"omitempty,required"`
	Description    *string  `json:"description,omitempty" validate:"omitempty,required"`
	Price          *float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	CompareAtPrice *float64 `json:"compare_at_price,omitempty" validate:"omitempty,gt=0"`
	CategoryID     *uint    `json:"category_id" validate:"omitempty,required"`
	Category       *string  `json:"category" validate:"omitempty"`
}
