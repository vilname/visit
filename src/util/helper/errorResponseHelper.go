// Package helper маленькие методы которые могут применяться по всему проекту
package helper

import (
	"net/http"
	"visit/src/util/constant"

	"github.com/gin-gonic/gin"
)

// ErrorResponse структура ответа об ошибке
type ErrorResponse struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"errorMessage"`
}

// ErrorResponseMethod обработка ошибки
func ErrorResponseMethod(ctx *gin.Context, err error) {
	errorResponse := ErrorResponse{}
	errorResponse.Message = findErrorType(err)
	errorResponse.ErrorMessage = err.Error()

	ctx.JSON(http.StatusInternalServerError, errorResponse)
}

func findErrorType(err error) string {
	var errorType constant.ErrorType

	switch err.Error() {
	case string(constant.MaxAttemptGenerateCode):
		errorType = constant.MaxAttempt
	default:
		errorType = constant.DateError
	}

	return string(errorType)
}
