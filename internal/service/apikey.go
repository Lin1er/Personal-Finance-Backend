// Package service contains the business logic for the application.
package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"personal-finance-backend/internal/domain"
	"personal-finance-backend/internal/repository"
)

type ApiKeyService struct {
	repo *repository.ApiKeyRepository
}

func NewApiKeyService(repo *repository.ApiKeyRepository) *ApiKeyService {
	return &ApiKeyService{repo: repo}
}

// generateKey creates a cryptographically secure random 64-char hex key.
func generateKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *ApiKeyService) Create(ctx context.Context, req domain.CreateApiKeyRequest) (*domain.ApiKey, error) {
	key, err := generateKey()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, req.Name, key)
}

func (s *ApiKeyService) GetAll(ctx context.Context) ([]domain.ApiKey, error) {
	return s.repo.GetAll(ctx)
}

func (s *ApiKeyService) GetByID(ctx context.Context, id string) (*domain.ApiKey, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ApiKeyService) Update(ctx context.Context, id string, req domain.UpdateApiKeyRequest) (*domain.ApiKey, error) {
	return s.repo.Update(ctx, id, req)
}

func (s *ApiKeyService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *ApiKeyService) ValidateKey(ctx context.Context, key string) (*domain.ApiKey, error) {
	return s.repo.ValidateKey(ctx, key)
}
