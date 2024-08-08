package utils

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"xendit-exercises/app/models"
)

var DefaultStatusText = map[int]string{
	http.StatusInternalServerError: "Server malfunction, please try again later",
	http.StatusNotFound:            "Data not found",
	http.StatusBadRequest:          "There is an error in the request data, please check again",
}

func WriteError(error error, errorCode int, message string) *models.ErrorLog {
	if pc, file, line, ok := runtime.Caller(1); ok {
		file = file[strings.LastIndex(file, "/")+1:]
		funcName := runtime.FuncForPC(pc).Name()
		output := &models.ErrorLog{
			StatusCode: errorCode,
			Error:      error,
		}

		output.SystemMessage = error.Error()
		if message == "" {
			output.Message = DefaultStatusText[errorCode]
			if output.Message == "" {
				output.Message = http.StatusText(errorCode)
			}
		} else {
			output.Message = message
		}

		if errorCode == http.StatusInternalServerError {
			output.Line = fmt.Sprintf("%d", line)
			output.Filename = file
			output.Function = funcName
		}
		return output
	}

	return nil
}
