package config

import (
	"net/http"
)

func ServiceCodeToHttpStatus(code ServiceCode) int {
	switch code {
	case CodeOK:
		return http.StatusOK
	case CodeBadRequest, CodeEmptyOrder:
		return http.StatusBadRequest
	case CodeUnprocessableEntity, CodeDatabaseError:
		return http.StatusUnprocessableEntity
	case CodeNotFound:
		return http.StatusNotFound
	case CodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func DBErrToServiceCode(err error) ServiceCode {
	switch err {
	case nil:
		return CodeOK
	case ErrRecordNotFound:
		return CodeNotFound
	default:
		return CodeDatabaseError
	}
}
