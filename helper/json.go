package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	// (1) Decode request
	decoder := json.NewDecoder(request.Body)
	// (2) Decode request to category request struct
	err := decoder.Decode(result)
	// (4) If error, handle with helper
	PanicErr(err)
}
