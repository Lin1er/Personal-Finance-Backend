package domain

import "time"

type Transaction struct {
	ID           string    `json:"id"`
	Type         string    `json:"type"`
	CategoryID   string    `json:"category_id"`
	CategoryName string    `json:"category_name,omitempty"` // joined from categories
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currency"`
	Description  *string   `json:"description,omitempty"`
	Status       string    `json:"status"`
	Date         string    `json:"date"` // YYYY-MM-DD
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateTransactionRequest struct {
	Type        string  `json:"type" binding:"required,oneof=income expense transfer"`
	CategoryID  string  `json:"category_id" binding:"required,uuid"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Currency    string  `json:"currency" binding:"omitempty,len=3"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status" binding:"omitempty,oneof=pending completed cancelled"`
	Date        string  `json:"date" binding:"omitempty"` // YYYY-MM-DD
}

type UpdateTransactionRequest struct {
	Type        *string  `json:"type,omitempty" binding:"omitempty,oneof=income expense transfer"`
	CategoryID  *string  `json:"category_id,omitempty" binding:"omitempty,uuid"`
	Amount      *float64 `json:"amount,omitempty" binding:"omitempty,gt=0"`
	Currency    *string  `json:"currency,omitempty" binding:"omitempty,len=3"`
	Description *string  `json:"description,omitempty"`
	Status      *string  `json:"status,omitempty" binding:"omitempty,oneof=pending completed cancelled"`
	Date        *string  `json:"date,omitempty"`
}

type TransactionFilter struct {
	Type       string `form:"type"`
	CategoryID string `form:"category_id"`
	Status     string `form:"status"`
	DateFrom   string `form:"date_from"` // YYYY-MM-DD
	DateTo     string `form:"date_to"`   // YYYY-MM-DD
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
}
