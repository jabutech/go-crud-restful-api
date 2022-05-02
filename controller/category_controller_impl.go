package controller

import (
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"go-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService // Use category service
}

func NewCategoryController(categorySevice service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categorySevice,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// (1) Create variable with value web.CategoryCreateRequest
	categoryCreateRequest := web.CategoryCreateRequest{}
	// (2) Decode with helper ReadFromRequestBody
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	// (3) Create new category use service Create
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	// (4) If success, create response with helper web response
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	// (5) Encode response with helper WriteToResponseBody
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// (1) Create variable with value web.CategoryUpdateRequest{
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	// (3) Decode with helper ReadFromRequestBody
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	// (4) Get parameter id
	categoryId := params.ByName("categoryId")
	// (5) Convert to string
	id, err := strconv.Atoi(categoryId)
	// (6) If error, handle with helper
	helper.PanicErr(err)

	// (7) Parse parameter id to categoryUpdateRequest
	categoryUpdateRequest.Id = id

	// (8) Update category use service Update
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	// (9) If success, create response with helper web response
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	// (10) Encode response with helper WriteToResponseBody
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// (1) Get parameter id
	categoryId := params.ByName("categoryId")
	// (2) Convert to string
	id, err := strconv.Atoi(categoryId)
	// (3) If error, handle with helper
	helper.PanicErr(err)

	// (5) Delete category use service Delete
	controller.CategoryService.Delete(request.Context(), id)

	// (6) If success, create response with helper web response
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	// (7) Encode response with helper WriteToResponseBody
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// (1) Get parameter id
	categoryId := params.ByName("categoryId")
	// (2) Convert to string
	id, err := strconv.Atoi(categoryId)
	// (3) If error, handle with helper
	helper.PanicErr(err)

	// (5) FindById category use service FindById
	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	// (6) If success, create response with helper web response
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	// (7) Encode response with helper WriteToResponseBody
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// (1) Delete category use service Delete
	webResponses := controller.CategoryService.FindAll(request.Context())

	// (6) If success, create response with helper web response
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   webResponses,
	}

	// (7) Encode response with helper WriteToResponseBody
	helper.WriteToResponseBody(writer, webResponse)

}
