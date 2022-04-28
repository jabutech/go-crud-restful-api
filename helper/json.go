package helper

import (
	"encoding/json"
	"net/http"
)

// Function for handle decode request body
func ReadFromRequestBody(request *http.Request, result interface{}) {
	// (1) Decode request
	decoder := json.NewDecoder(request.Body)
	// (2) Decode request to category request struct
	err := decoder.Decode(result)
	// (4) If error, handle with helper
	PanicErr(err)
}

// function for handle encode response body
func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	// (1) Add header
	writer.Header().Add("Content-Type", "application/json")
	// (8) Create encode
	encoder := json.NewEncoder(writer)
	// (9) Encode web response
	err := encoder.Encode(response)
	// (10) If error, handle with helper
	PanicErr(err)
}
