package responses

import (
	"errors"
	"net/http"

	domainerrors "real-estate-agency-onion/internal/domain/errors"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func MapDomainError(err error) (int, ErrorResponse) {
	switch {
	case errors.Is(err, domainerrors.ErrNotFound):
		return http.StatusNotFound, ErrorResponse{Error: "not_found", Message: err.Error()}
	case errors.Is(err, domainerrors.ErrAlreadyExists):
		return http.StatusConflict, ErrorResponse{Error: "conflict", Message: err.Error()}
	case errors.Is(err, domainerrors.ErrInvalidInput):
		return http.StatusBadRequest, ErrorResponse{Error: "bad_request", Message: err.Error()}
	case errors.Is(err, domainerrors.ErrForbidden):
		return http.StatusForbidden, ErrorResponse{Error: "forbidden", Message: err.Error()}
	default:
		return http.StatusInternalServerError, ErrorResponse{Error: "internal", Message: "internal server error"}
	}
}