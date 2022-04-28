package controller

import (
	"encoding/json"
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

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// (1) Decode request
	decoder := json.NewDecoder(request.Body)
	// (2) Create variable with value web.CategoryCreateRequest{
	categoryCreateRequest := web.CategoryCreateRequest{}
	// (3) Decode request to category request struct
	err := decoder.Decode(&categoryCreateRequest)
	// (4) If error, handle with helper
	helper.PanicErr(err)

	// (5) Create new category use service Create
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	// (6) If success, create response with helper web response
	webWebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	// (7) Add header
	writer.Header().Add("Content-Type", "application/json")
	// (8) Create encode
	encoder := json.NewEncoder(writer)
	// (9) Encode web response
	err = encoder.Encode(webWebResponse)
	// (10) If error, handle with helper
	helper.PanicErr(err)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// (1) Decode request
	decoder := json.NewDecoder(request.Body)
	// (2) Create variable with value web.CategoryUpdateRequest{
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	// (3) Decode request to category request struct
	err := decoder.Decode(&categoryUpdateRequest)
	// (4) If error, handle with helper
	helper.PanicErr(err)

	// (5) Get parameter id
	categoryId := params.ByName("categoryId")
	// (6) Convert to string
	id, err := strconv.Atoi(categoryId)
	// (7) If error, handle with helper
	helper.PanicErr(err)

	// (8) Parse parameter id to categoryUpdateRequest
	categoryUpdateRequest.Id = id

	// (5) Update category use service Update
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	// (6) If success, create response with helper web response
	webWebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	// (7) Add header
	writer.Header().Add("Content-Type", "application/json")
	// (8) Create encode
	encoder := json.NewEncoder(writer)
	// (9) Encode web response
	err = encoder.Encode(webWebResponse)
	// (10) If error, handle with helper
	helper.PanicErr(err)

}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// (1) Get parameter id
	categoryId := params.ByName("categoryId")
	// (1) Convert to string
	id, err := strconv.Atoi(categoryId)
	// (1) If error, handle with helper
	helper.PanicErr(err)

	// (5) Delete category use service Delete
	controller.CategoryService.Delete(request.Context(), id)

	// (6) If success, create response with helper web response
	webWebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	// (7) Add header
	writer.Header().Add("Content-Type", "application/json")
	// (8) Create encode
	encoder := json.NewEncoder(writer)
	// (9) Encode web response
	err = encoder.Encode(webWebResponse)
	// (10) If error, handle with helper
	helper.PanicErr(err)

}
