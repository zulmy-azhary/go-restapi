package dto

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"omitempty,min=3"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"omitempty,gt=0"`
	Stock       int     `json:"stock" validate:"omitempty,gte=0"`
}

type ProductQuery struct {
	Page   int    `query:"page"`
	PerPage  int    `query:"perPage"`
	Search string `query:"search"`
}