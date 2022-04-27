package helper

import (
	"go-restful-api/model/domain"
	"go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	// (3) Create new variable
	var categoryResponses []web.CategoryResponse

	// (4) Loop all data categories and append to variable categoryResponses
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
