package repository

import (
	"context"
	"database/sql"
	"go-restful-api/model/domain"
)

// Contract for repository category
type CategoryRespository interface {
	// Contract function Save for insert data
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	// Contract function Update for update data
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	// Contract function Delete for delete data
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category) string
	// Contract function FindId for find data based on id
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	// Contract function FindAll for find all data
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
