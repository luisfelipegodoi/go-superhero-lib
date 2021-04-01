package error

import (
	"errors"
	"net/http"

	"github.com/luisfelipegodoi/go-superhero-lib/src/context/general/logger"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")

	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")

	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")

	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")

	// ErrBadRequest bad request ocurred when requisition contains errors
	ErrBadRequest = errors.New("your request was recused because contains errors")
)

// GetStatusErrorCode - Check error type and return approprieted
func GetStatusErrorCode(err error) int {

	logger.Error(err)

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrBadParamInput:
		return http.StatusBadRequest
	case ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
