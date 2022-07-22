package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpError struct {
	code    int
	message string
	cause   error
}

func (he HttpError) Error() string {
	return he.cause.Error()
}

func (he HttpError) Code() int {
	return he.code
}

func (he HttpError) Cause() error {
	return he.cause
}

func (he HttpError) Message() string {
	return he.message
}

func (he HttpError) toJson() gin.H {
	return gin.H{"statusCode": he.Code(), "message": he.Message()}
}

func (he HttpError) Is(target error) bool {
	if he == target {
		return true
	}

	if x, ok := target.(interface{ Code() int }); ok {
		if x.Code() == he.Code() {
			return true
		}
	}

	return false
}

func NotFoundError(resourceName string, cause error) error {
	return HttpError{
		code:    http.StatusNotFound,
		message: fmt.Sprintf("resource '%s' not found", resourceName),
		cause:   cause,
	}
}

func InternalServerError(cause error) error {
	return HttpError{
		code:    http.StatusInternalServerError,
		message: "InternalServerError",
		cause:   cause,
	}
}
