package api

import (
	"github.com/HamidSajjadi/ushort/internal"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var (
	notFoundError       = HttpError{code: http.StatusNotFound, message: "resrouce not found"}
	internalServerError = HttpError{code: http.StatusInternalServerError, message: "internal server error"}
	conflictError       = HttpError{code: http.StatusConflict, message: "resource already exists"}
)

var errorToHttpErrorMap = map[error]HttpError{
	internal.NotFoundErr: notFoundError,
	internal.ConflictErr: conflictError,
}

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

func createHttpErrorFromGenericError(err error) *HttpError {
	httpErr := HttpError{}
	if e, ok := err.(interface{ Code() int }); ok {
		httpErr.code = e.Code()
	} else {
		httpErr.code = http.StatusInternalServerError
	}
	if e, ok := err.(interface{ Message() string }); ok {
		httpErr.message = e.Message()
	} else {
		httpErr.message = ""
	}
	if e, ok := err.(interface{ Cause() error }); ok {
		httpErr.cause = e.Cause()
	}
	return &httpErr
}

func toHttpError(err error) *HttpError {
	httpErr := HttpError{}
	if u, ok := err.(HttpError); ok {
		httpErr = u
	} else if u, ok = errorToHttpErrorMap[err]; ok {
		httpErr = u
	} else {
		httpErr = *createHttpErrorFromGenericError(&httpErr)
	}
	httpErr.cause = err

	return &httpErr
}

func ErrorHandler(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			e := c.Errors[0]
			httpError := toHttpError(e.Err)

			logger.Errorw("error in API",
				"httpCode", httpError.Code(),
				"cause", httpError.Cause(),
				"message", httpError.Message(),
			)

			c.JSON(httpError.Code(), httpError.toJson())
		}
	}

}
