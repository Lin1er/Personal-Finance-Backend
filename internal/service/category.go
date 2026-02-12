package service

import (
	"context"

	"personal-finance-backend/internal/domain"
	"personal-finance-backend/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(ctx context.Context, req domain.CreateCategoryRequest) (*domain.Category, error) {
	return s.repo.Create(ctx, req)
}

func (s *CategoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *CategoryService) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CategoryService) Update(ctx context.Context, id string, req domain.UpdateCategoryRequest) (*domain.Category, error) {
	return s.repo.Update(ctx, id, req)
}

func (s *CategoryService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
