package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jabutech/go-crud-restful-api/helper"
	"github.com/jabutech/go-crud-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoriRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

// Function Save with follow the contract category repository
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// (1) Create sql query
	SQL := "insert into category(name) values (?)"

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

// Function Update with follow the contract category repository
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// (1) Create sql query
	SQL := "update category set name = ? where id = ?"

	// (2) Create context
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)

	// (3) If error, handle with helper error
	helper.PanicErr(err)

	// (4) If success, return category
	return category
}

// Function Delete with follow the contract category repository
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) string {
	// (1) Create sql query
	SQL := "delete from category where id = ?"

	// (2) Create context
	_, err := tx.ExecContext(ctx, SQL, category.Id)

	// (3) If error handle with helper error
	helper.PanicErr(err)

	// (4) Return info success deleted
	return "Delete success."
}

// Function Find data by id with follow the contract category repository
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	// (1) Create sql query
	SQL := "select id, name from category where id = ?"

	// (2) Create query context
	rows, err := tx.QueryContext(ctx, SQL, categoryId)

	// (3) If error, handle with helper error
	helper.PanicErr(err)

	// (4) Close rows after use
	defer rows.Close()

	// (5) Create variable category with value struct domain category
	category := domain.Category{}

	// (6) If category is available
	if rows.Next() {
		// (1) Get data category
		err := rows.Scan(&category.Id, &category.Name)

		// (2) If error, handle with helper error
		helper.PanicErr(err)

		// (3) If no, return category with error nil
		return category, nil
	} else {
		// If category is empty, return category and send info error
		return category, errors.New("category is not found")
	}
}

// Function Find all data with follow the contract category repository
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	// (1) Create sql query
	SQL := "select id, name from category"

	// (2) Create query context
	rows, err := tx.QueryContext(ctx, SQL)

	// (3) If error, handle with helper error
	helper.PanicErr(err)

	// (4) Close rows after use
	defer rows.Close()

	// (5) Create var category with value slice domain category
	var categories []domain.Category

	// (6) If category is available
	for rows.Next() {
		// (1) Create var category with strucy domain category
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)

		// (2) If error, handle with helper error
		helper.PanicErr(err)

		// (3) If no error, insert all data to var categories
		categories = append(categories, category)
	}

	// (7) return all data category
	return categories
}
