package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/pohanson/cvwo-forum/internal/model"
	"github.com/pohanson/cvwo-forum/internal/repository"
)

func writeDecodingError(err error, w http.ResponseWriter) {
	if err != nil {
		var syntaxError *json.SyntaxError
		switch {
		case errors.As(err, &syntaxError):
			http.Error(w, fmt.Sprint("Error decoding json: ", err.Error()), http.StatusBadRequest)
		default:
			http.Error(w, "Error decoding json", http.StatusBadRequest)
		}
	}
}

func writeValidationError(err error, w http.ResponseWriter) {
	if err != nil {
		var maxLengthError *model.MaxLengthViolationErr
		var missingFieldError *model.MissingFieldErr
		var invalidFK *model.InvalidFK
		switch {
		case errors.As(err, &maxLengthError):
			fallthrough
		case errors.As(err, &missingFieldError):
			fallthrough
		case errors.As(err, &invalidFK):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "Invalid Json Data", http.StatusBadRequest)
		}
	}

}

func writeDbInsertError(err error, w http.ResponseWriter) {
	var uniqueViolationErr *repository.UniqueViolationErr
	switch {
	case errors.As(err, &uniqueViolationErr):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, "Unknown Error", http.StatusBadRequest)
	}
}
