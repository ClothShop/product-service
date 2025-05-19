package product

type Update struct {
	ID             uint     `json:"id"`
	Name           *string  `json:"name,omitempty" validate:"omitempty,required"`
	Description    *string  `json:"description,omitempty" validate:"omitempty,required"`
	Price          *float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	CompareAtPrice *float64 `json:"compare_at_price" validate:"omitempty"`
	CategoryID     *uint    `json:"category_id" validate:"omitempty,required"`
	Category       *string  `json:"category" validate:"omitempty"`
}
