package category

type Update struct {
	Name string `json:"name" validate:"omitempty;required"`
}
