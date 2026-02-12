package domain

import "time"

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"` // "income" or "expense"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
	Type string `json:"type" binding:"required,oneof=income expense"`
}

type UpdateCategoryRequest struct {
	Name *string `json:"name,omitempty" binding:"omitempty,min=1,max=50"`
	Type *string `json:"type,omitempty" binding:"omitempty,oneof=income expense"`
}
