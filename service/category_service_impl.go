package service

import (
	"context"
	"database/sql"
	"go-restful-api/exception"
	"go-restful-api/helper"
	"go-restful-api/model/domain"
	"go-restful-api/model/web"
	"go-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository // Use repository
	DB                 *sql.DB                       // Use Sql driver
	Validate           *validator.Validate           // Use validator
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// Function service for proses create new category
func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// (1) Run validate before create data
	err := service.Validate.Struct(request)
	// (2) If error, handle with helper
	helper.PanicErr(err)

	// (3) Create transactional database
	tx, err := service.DB.Begin()
	// (4) Handle if create transaction error
	helper.PanicErr(err)
	// (5) Run this process in the end all operation with defer, and check process transaction Commit or Rollback transaction
	defer helper.CommitOrRollback(tx)

	// (6) Create new object category
	category := domain.Category{
		// Set name from request
		Name: request.Name,
	}

	// (7) Save transaction with use Repository
	category = service.CategoryRepository.Save(ctx, tx, category)

	// (8) Return after success
	return helper.ToCategoryResponse(category)
}

// Function service for proses update category
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// (1) Run validate before create data
	err := service.Validate.Struct(request)
	// (2) If error, handle with helper
	helper.PanicErr(err)

	// (3) Create transactional database
	tx, err := service.DB.Begin()
	// (4) Handle if create transaction error
	helper.PanicErr(err)
	// (5) Run this process in the end all operation with defer, and check process transaction Commit or Rollback transaction
	defer helper.CommitOrRollback(tx)

	// (6) Find category in dataabase
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)

	// (7) If error / category not found handle error with exception not found
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// (8) If no error, set request name to object category
	category.Name = request.Name

	// (9) Update category with use Repository
	category = service.CategoryRepository.Update(ctx, tx, category)

	// (10) Return response with helper
	return helper.ToCategoryResponse(category)
}

// Function service for process delete category
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// (1) Create transactional database
	tx, err := service.DB.Begin()
	// (2) Handle if create transaction error
	helper.PanicErr(err)
	// (3) Run this process in the end all operation with defer, and check process transaction Commit or Rollback transaction
	defer helper.CommitOrRollback(tx)

	// (2) Find category by id with use Repository
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	//  (3) If error / category not found handle error with exception
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// (4) If no error, Delete category
	service.CategoryRepository.Delete(ctx, tx, category)
}

// Function service for process delete category
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	// (1) Create transactional database
	tx, err := service.DB.Begin()
	// (2) Handle if create transaction error
	helper.PanicErr(err)
	// (3) Run this process in the end all operation with defer, and check process transaction Commit or Rollback transaction
	defer helper.CommitOrRollback(tx)

	// (2) Find category by id with use Repository
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	//  (3) If error / category not found handle error with exception
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// (4) If no error, Return category
	return helper.ToCategoryResponse(category)
}

// Function service for process delete category
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// (1) Create transactional database
	tx, err := service.DB.Begin()
	// (2) Handle if create transaction error
	helper.PanicErr(err)
	// (3) Run this process in the end all operation with defer, and check process transaction Commit or Rollback transaction
	defer helper.CommitOrRollback(tx)

	// (2) Get all categories
	categories := service.CategoryRepository.FindAll(ctx, tx)

	// (3)  Return with helper ToCategoryResponses
	return helper.ToCategoryResponses(categories)
}
