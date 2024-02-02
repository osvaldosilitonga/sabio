package dto

type ProductReq struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	Stock int     `json:"stock" validate:"required"`
}

type UpdateReq struct {
	ID    int     `json:"id" validate:"required"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitemty"`
	Stock int     `json:"stock,omitempty"`
}
