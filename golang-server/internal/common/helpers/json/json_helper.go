package json_helper

import (
	"encoding/json"
	"net/http"
	"rz-server/internal/common/errors/application_error"
	"rz-server/internal/common/interfaces"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Key     string `json:"code"`
	Context any    `json:"context"`
}

func RespondJsonError(err interfaces.ApplicationError, w http.ResponseWriter) {
	statusError := map[application_error.ErrorType]int{
		application_error.BusinessError:   http.StatusUnprocessableEntity,
		application_error.ValidationError: http.StatusUnprocessableEntity,
		application_error.InvalidInput:    http.StatusBadRequest,
		application_error.NotAllowed:      http.StatusForbidden,
		application_error.NotFound:        http.StatusNotFound,
		application_error.InternalError:   http.StatusServiceUnavailable,
	}

	data := application_error.ErrorType(err.GetErrorType())

	statusCode := statusError[data]

	w.WriteHeader(statusCode)

	jsonResponse := ErrorResponse{
		Message: err.GetMessage(),
		Key:     err.GetKey(),
		Context: err.GetContext(),
	}

	jsonResponseBytes, _ := json.Marshal(jsonResponse)
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponseBytes)
}

func RespondJsonResourceSuccess[T any](mapper interfaces.ResourceMapper[T], w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, _ := json.Marshal(mapper.ToResource())
	w.Write(jsonResponse)
}
