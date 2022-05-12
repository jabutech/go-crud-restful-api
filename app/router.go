package app

import (
	"net/http"

	"github.com/jabutech/go-crud-restful-api/controller"
	"github.com/jabutech/go-crud-restful-api/exception"
	"github.com/jabutech/go-crud-restful-api/helper"
	"github.com/jabutech/go-crud-restful-api/model/web"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	// Use http router
	router := httprouter.New()

	// Endpoint
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "Server is Up.",
		}

		// (5) Encode response with helper WriteToResponseBody
		helper.WriteToResponseBody(w, webResponse)
	})
	// Get all categories
	router.GET("/api/categories", categoryController.FindAll)
	// Get category by id
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	// Create new category
	router.POST("/api/categories", categoryController.Create)
	// Update category by id
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	// Delete category by id
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Change PanicHandler to exception error hanlder
	router.PanicHandler = exception.ErrorHandler

	return router
}
