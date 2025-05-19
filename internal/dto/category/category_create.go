package category

type Create struct {
	Name string `json:"name" validate:"required"`
}
