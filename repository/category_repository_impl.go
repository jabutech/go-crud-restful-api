package repository

import (
	"context"
	"database/sql"
	"go-restful-api/helper"
	"go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

// Function Save with follow the contract category repository
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// (1) Create sql query
	SQL := "insert into customer(name) valus (?)"

	// (2) Create context
	result, err := tx.ExecContext(ctx, SQL, category.Name)

	// (3) If error handle error with helper error
	helper.PanicErr(err)

	// (4) If success, get last insert id
	id, err := result.LastInsertId()

	// (5) Handle if error
	helper.PanicErr(err)

	// (6) Set last insert id to category id and convert from type int64 to int
	category.Id = int(id)

	// (7) Return category
	return category
}
