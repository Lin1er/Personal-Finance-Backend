package service

import (
	"context"

	"personal-finance-backend/internal/domain"
	"personal-finance-backend/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(ctx context.Context, req domain.CreateTransactionRequest) (*domain.Transaction, error) {
	return s.repo.Create(ctx, req)
}

func (s *TransactionService) GetAll(ctx context.Context, filter domain.TransactionFilter) ([]domain.Transaction, int, error) {
	return s.repo.GetAll(ctx, filter)
}

func (s *TransactionService) GetByID(ctx context.Context, id string) (*domain.Transaction, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TransactionService) Update(ctx context.Context, id string, req domain.UpdateTransactionRequest) (*domain.Transaction, error) {
	return s.repo.Update(ctx, id, req)
}

func (s *TransactionService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
