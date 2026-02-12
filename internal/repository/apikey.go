// Package repository contains the database queries for the application.
package repository

import (
	"context"

	"personal-finance-backend/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ApiKeyRepository struct {
	db *pgxpool.Pool
}

func NewApiKeyRepository(db *pgxpool.Pool) *ApiKeyRepository {
	return &ApiKeyRepository{db: db}
}

func (r *ApiKeyRepository) Create(ctx context.Context, name, key string) (*domain.ApiKey, error) {
	var apiKey domain.ApiKey
	err := r.db.QueryRow(ctx,
		`INSERT INTO api_keys (name, key) VALUES ($1, $2)
		 RETURNING id, name, key, is_active, created_at, last_used_at`,
		name, key,
	).Scan(&apiKey.ID, &apiKey.Name, &apiKey.Key, &apiKey.IsActive, &apiKey.CreatedAt, &apiKey.LastUsedAt)
	if err != nil {
		return nil, err
	}
	return &apiKey, nil
}

func (r *ApiKeyRepository) GetAll(ctx context.Context) ([]domain.ApiKey, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, is_active, created_at, last_used_at
		 FROM api_keys ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keys []domain.ApiKey
	for rows.Next() {
		var k domain.ApiKey
		if err := rows.Scan(&k.ID, &k.Name, &k.IsActive, &k.CreatedAt, &k.LastUsedAt); err != nil {
			return nil, err
		}
		keys = append(keys, k)
	}
	return keys, nil
}

func (r *ApiKeyRepository) GetByID(ctx context.Context, id string) (*domain.ApiKey, error) {
	var k domain.ApiKey
	err := r.db.QueryRow(ctx,
		`SELECT id, name, is_active, created_at, last_used_at
		 FROM api_keys WHERE id = $1`, id,
	).Scan(&k.ID, &k.Name, &k.IsActive, &k.CreatedAt, &k.LastUsedAt)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *ApiKeyRepository) Update(ctx context.Context, id string, req domain.UpdateApiKeyRequest) (*domain.ApiKey, error) {
	var k domain.ApiKey

	// Build dynamic update — only update fields that are provided
	if req.Name != nil {
		if _, err := r.db.Exec(ctx, `UPDATE api_keys SET name = $1 WHERE id = $2`, *req.Name, id); err != nil {
			return nil, err
		}
	}
	if req.IsActive != nil {
		if _, err := r.db.Exec(ctx, `UPDATE api_keys SET is_active = $1 WHERE id = $2`, *req.IsActive, id); err != nil {
			return nil, err
		}
	}

	err := r.db.QueryRow(ctx,
		`SELECT id, name, is_active, created_at, last_used_at
		 FROM api_keys WHERE id = $1`, id,
	).Scan(&k.ID, &k.Name, &k.IsActive, &k.CreatedAt, &k.LastUsedAt)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *ApiKeyRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM api_keys WHERE id = $1`, id)
	return err
}

// ValidateKey checks if a key exists and is active. Updates last_used_at.
func (r *ApiKeyRepository) ValidateKey(ctx context.Context, key string) (*domain.ApiKey, error) {
	var k domain.ApiKey
	err := r.db.QueryRow(ctx,
		`UPDATE api_keys SET last_used_at = now()
		 WHERE key = $1 AND is_active = true
		 RETURNING id, name, is_active, created_at, last_used_at`,
		key,
	).Scan(&k.ID, &k.Name, &k.IsActive, &k.CreatedAt, &k.LastUsedAt)
	if err != nil {
		return nil, err
	}
	return &k, nil
}
