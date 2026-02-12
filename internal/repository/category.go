package repository

import (
	"context"

	"personal-finance-backend/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, req domain.CreateCategoryRequest) (*domain.Category, error) {
	var c domain.Category
	err := r.db.QueryRow(ctx,
		`INSERT INTO categories (name, type) VALUES ($1, $2)
		 RETURNING id, name, type, created_at, updated_at`,
		req.Name, req.Type,
	).Scan(&c.ID, &c.Name, &c.Type, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, type, created_at, updated_at
		 FROM categories ORDER BY type, name`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var c domain.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Type, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	var c domain.Category
	err := r.db.QueryRow(ctx,
		`SELECT id, name, type, created_at, updated_at
		 FROM categories WHERE id = $1`, id,
	).Scan(&c.ID, &c.Name, &c.Type, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id string, req domain.UpdateCategoryRequest) (*domain.Category, error) {
	if req.Name != nil {
		if _, err := r.db.Exec(ctx, `UPDATE categories SET name = $1, updated_at = now() WHERE id = $2`, *req.Name, id); err != nil {
			return nil, err
		}
	}
	if req.Type != nil {
		if _, err := r.db.Exec(ctx, `UPDATE categories SET type = $1, updated_at = now() WHERE id = $2`, *req.Type, id); err != nil {
			return nil, err
		}
	}
	return r.GetByID(ctx, id)
}

func (r *CategoryRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM categories WHERE id = $1`, id)
	return err
}
