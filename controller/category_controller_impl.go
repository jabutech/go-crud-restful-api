package controller

import (
	"encoding/json"
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"go-restful-api/service"
	"net/http"

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
