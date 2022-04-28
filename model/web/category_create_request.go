package web

// Struct for request create new data
type CategoryCreateRequest struct {
	Name string `validate:"required,max=200,min=1"`
}
