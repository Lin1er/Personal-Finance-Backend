// Package domain contains the domain models for the application.
package domain

import "time"

type ApiKey struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Key        string     `json:"key,omitempty"` // only shown on creation
	IsActive   bool       `json:"is_active"`
	CreatedAt  time.Time  `json:"created_at"`
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
}

type CreateApiKeyRequest struct {
	Name string `json:"name" binding:"required,min=1,max=100"`
}

type UpdateApiKeyRequest struct {
	Name     *string `json:"name,omitempty" binding:"omitempty,min=1,max=100"`
	IsActive *bool   `json:"is_active,omitempty"`
}
