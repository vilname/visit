// Package restadmin контроллер для админской части докторов
package restadmin

import (
	"net/http"
	"visit/src/model"

	"github.com/gin-gonic/gin"
)

// Doctor godoc
// @Tags Доктор
// @Summary Список промокодов
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Security OAuth2Implicit
//
// @Failure	500	{object} helper.ErrorResponse "Другие ошибки"
// @Router /admin/doctor [get]
func Doctor(ctx *gin.Context) {
	doctor := model.Doctor{
		Name: "Admin",
	}

	ctx.JSON(http.StatusOK, doctor)
}
