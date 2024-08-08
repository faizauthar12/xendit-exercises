package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponseBody(data interface{}) []byte {
	responseJson, err := json.Marshal(data)

	if err != nil {
		errorLog := WriteError(err, http.StatusInternalServerError, "")
		errorLogJson, _ := json.Marshal(errorLog)
		return errorLogJson
	}
	return responseJson
}
