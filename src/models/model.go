package models

type (
	Order struct {
		Title string  `json:"title" validator:"required,max=64"`
		Price float64 `json:"price" validator:"required"`
	}

	ErrResponse struct {
		Message string `json:"message"`
	}
)
