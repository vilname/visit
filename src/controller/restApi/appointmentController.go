// Package restapi контроллеры публичной части
package restapi

import (
	"net/http"
	"visit/src/model"
	"visit/src/service"
	"visit/src/util/helper"

	"github.com/gin-gonic/gin"
)

// CreateAppointment godoc
// @Tags Запись на приём
// @Summary Создание записи на приём к врачу
// @Description Создание новой записи на приём к выбранному врачу
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.AppointmentCreateRequest true "Данные для записи"
// @Success 201 {object} model.AppointmentResponse "Созданная запись"
// @Failure 400 {object} helper.ErrorResponse "Ошибка валидации"
// @Failure 401 {object} helper.ErrorResponse "Не авторизован"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/appointments [post]
func CreateAppointment(ctx *gin.Context) {
	userID, exists := ctx.Get("userUuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, helper.ErrorResponse{
			Message:      "AUTH_ERROR",
			ErrorMessage: "пользователь не авторизован",
		})
		return
	}

	var req model.AppointmentCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Message:      "VALIDATION_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	appointment, err := service.CreateAppointment(userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Message:      "APPOINTMENT_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, appointment)
}

// GetAppointments godoc
// @Tags Запись на приём
// @Summary Список записей пользователя
// @Description Получение списка всех записей на приём текущего пользователя
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} model.AppointmentResponse "Список записей"
// @Failure 401 {object} helper.ErrorResponse "Не авторизован"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/appointments [get]
func GetAppointments(ctx *gin.Context) {
	userID, exists := ctx.Get("userUuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, helper.ErrorResponse{
			Message:      "AUTH_ERROR",
			ErrorMessage: "пользователь не авторизован",
		})
		return
	}

	appointments, err := service.GetUserAppointments(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Message:      "APPOINTMENT_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, appointments)
}

// CancelAppointment godoc
// @Tags Запись на приём
// @Summary Отмена записи на приём
// @Description Отмена записи на приём по её идентификатору
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "ID записи"
// @Success 200 {object} map[string]string "Запись отменена"
// @Failure 400 {object} helper.ErrorResponse "Ошибка при отмене"
// @Failure 401 {object} helper.ErrorResponse "Не авторизован"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/appointments/{id} [delete]
func CancelAppointment(ctx *gin.Context) {
	userID, exists := ctx.Get("userUuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, helper.ErrorResponse{
			Message:      "AUTH_ERROR",
			ErrorMessage: "пользователь не авторизован",
		})
		return
	}

	appointmentID := ctx.Param("id")
	if appointmentID == "" {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Message:      "VALIDATION_ERROR",
			ErrorMessage: "не указан ID записи",
		})
		return
	}

	err := service.CancelAppointment(appointmentID, userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Message:      "CANCEL_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "запись успешно отменена"})
}
