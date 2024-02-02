package entity

type Order struct {
	ID         int     `json:"id,omitempty"`
	CustomerID int     `json:"customer_id"`
	ProductID  int     `json:"product_id"`
	Qty        int     `json:"quantity"`
	Total      float64 `json:"total"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
