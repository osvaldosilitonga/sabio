package dto

type CustomerReq struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UpdateReq struct {
	ID    int    `json:"id" validate:"required"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitemty"`
}
