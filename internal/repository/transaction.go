package repository

import (
	"context"
	"fmt"
	"strings"

	"personal-finance-backend/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, req domain.CreateTransactionRequest) (*domain.Transaction, error) {
	// Set defaults
	if req.Currency == "" {
		req.Currency = "IDR"
	}
	if req.Status == "" {
		req.Status = "completed"
	}
	if req.Date == "" {
		req.Date = "now()"
	}

	var t domain.Transaction
	err := r.db.QueryRow(ctx,
		`INSERT INTO transactions (type, category_id, amount, currency, description, status, date)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 RETURNING id, type, category_id, amount, currency, description, status, date::text, created_at, updated_at`,
		req.Type, req.CategoryID, req.Amount, req.Currency, req.Description, req.Status, req.Date,
	).Scan(&t.ID, &t.Type, &t.CategoryID, &t.Amount, &t.Currency, &t.Description, &t.Status, &t.Date, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TransactionRepository) GetAll(ctx context.Context, filter domain.TransactionFilter) ([]domain.Transaction, int, error) {
	// Set defaults
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 || filter.Limit > 100 {
		filter.Limit = 20
	}
	offset := (filter.Page - 1) * filter.Limit

	// Build WHERE clause dynamically
	var conditions []string
	var args []interface{}
	argIdx := 1

	if filter.Type != "" {
		conditions = append(conditions, fmt.Sprintf("t.type = $%d", argIdx))
		args = append(args, filter.Type)
		argIdx++
	}
	if filter.CategoryID != "" {
		conditions = append(conditions, fmt.Sprintf("t.category_id = $%d", argIdx))
		args = append(args, filter.CategoryID)
		argIdx++
	}
	if filter.Status != "" {
		conditions = append(conditions, fmt.Sprintf("t.status = $%d", argIdx))
		args = append(args, filter.Status)
		argIdx++
	}
	if filter.DateFrom != "" {
		conditions = append(conditions, fmt.Sprintf("t.date >= $%d", argIdx))
		args = append(args, filter.DateFrom)
		argIdx++
	}
	if filter.DateTo != "" {
		conditions = append(conditions, fmt.Sprintf("t.date <= $%d", argIdx))
		args = append(args, filter.DateTo)
		argIdx++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Count total
	var total int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM transactions t %s`, whereClause)
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Fetch data with JOIN to get category name
	query := fmt.Sprintf(`
		SELECT t.id, t.type, t.category_id, c.name, t.amount, t.currency,
		       t.description, t.status, t.date::text, t.created_at, t.updated_at
		FROM transactions t
		JOIN categories c ON c.id = t.category_id
		%s
		ORDER BY t.date DESC, t.created_at DESC
		LIMIT $%d OFFSET $%d`,
		whereClause, argIdx, argIdx+1,
	)
	args = append(args, filter.Limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		var t domain.Transaction
		if err := rows.Scan(
			&t.ID, &t.Type, &t.CategoryID, &t.CategoryName,
			&t.Amount, &t.Currency, &t.Description, &t.Status,
			&t.Date, &t.CreatedAt, &t.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		transactions = append(transactions, t)
	}
	return transactions, total, nil
}

func (r *TransactionRepository) GetByID(ctx context.Context, id string) (*domain.Transaction, error) {
	var t domain.Transaction
	err := r.db.QueryRow(ctx,
		`SELECT t.id, t.type, t.category_id, c.name, t.amount, t.currency,
		        t.description, t.status, t.date::text, t.created_at, t.updated_at
		 FROM transactions t
		 JOIN categories c ON c.id = t.category_id
		 WHERE t.id = $1`, id,
	).Scan(&t.ID, &t.Type, &t.CategoryID, &t.CategoryName, &t.Amount, &t.Currency,
		&t.Description, &t.Status, &t.Date, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TransactionRepository) Update(ctx context.Context, id string, req domain.UpdateTransactionRequest) (*domain.Transaction, error) {
	if req.Type != nil {
		if _, err := r.db.Exec(ctx, `UPDATE transactions SET type = $1, updated_at = now() WHERE id = $2`, *req.Type, id); err != nil {
			return nil, err
		}
	}
	if req.CategoryID != nil {
		if _, err := r.db.Exec(ctx, `UPDATE transactions SET category_id = $1, updated_at = now() WHERE id = $2`, *req.CategoryID, id); err != nil {
			return nil, err
		}
	}
	if req.Amount != nil {
		if _, err := r.db.Exec(ctx, `UPDATE transactions SET amount = $1, updated_at = now() WHERE id = $2`, *req.Amount, id); err != nil {
			return nil, err
		}
	}
	if req.Currency != nil {
		if _, err := r.db.Exec(ctx, `UPDATE transactions SET currency = $1, updated_at = now() WHERE id = $2`, *req.Currency, id); err != nil {
			return nil, err
		}
	}
	if req.Description != nil {
		if _, err := r.db.Exec(ctx, `UPDATE transactions SET description = $1, updated_at = now() WHERE id = $2`, *req.Description, id); err != nil {
			return nil, err
		}
	}
	if req.Status != nil {
		if _, err := r.db.Exec(ctx, `UPDATE transactions SET status = $1, updated_at = now() WHERE id = $2`, *req.Status, id); err != nil {
			return nil, err
		}
	}
	if req.Date != nil {
		if _, err := r.db.Exec(ctx, `UPDATE transactions SET date = $1, updated_at = now() WHERE id = $2`, *req.Date, id); err != nil {
			return nil, err
		}
	}
	return r.GetByID(ctx, id)
}

func (r *TransactionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM transactions WHERE id = $1`, id)
	return err
}
