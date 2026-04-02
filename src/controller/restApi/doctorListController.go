// Package restapi контроллеры публичной части
package restapi

import (
	"net/http"
	"visit/src/service"
	"visit/src/util/helper"

	"github.com/gin-gonic/gin"
)

// GetDoctors godoc
// @Tags Врачи
// @Summary Список врачей
// @Description Получение списка всех врачей медицинской организации. Можно фильтровать по специализации
// @Accept json
// @Produce json
// @Param specialization query string false "Фильтр по специализации"
// @Success 200 {object} model.DoctorListResponse "Список врачей"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/doctors [get]
func GetDoctors(ctx *gin.Context) {
	specialization := ctx.Query("specialization")

	var err error
	if specialization != "" {
		result, err := service.GetDoctorsBySpecialization(specialization)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{
				Message:      "DOCTOR_ERROR",
				ErrorMessage: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	}

	result, err := service.GetAllDoctors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Message:      "DOCTOR_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetDoctorByID godoc
// @Tags Врачи
// @Summary Получение врача по ID
// @Description Получение подробной информации о враче по его идентификатору
// @Accept json
// @Produce json
// @Param id path string true "ID врача"
// @Success 200 {object} model.DoctorResponse "Информация о враче"
// @Failure 404 {object} helper.ErrorResponse "Врач не найден"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/doctors/{id} [get]
func GetDoctorByID(ctx *gin.Context) {
	doctorID := ctx.Param("id")
	if doctorID == "" {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Message:      "VALIDATION_ERROR",
			ErrorMessage: "не указан ID врача",
		})
		return
	}

	doctor, err := service.GetDoctorByID(doctorID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.ErrorResponse{
			Message:      "NOT_FOUND",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, doctor)
}
