// Package restapi контроллеры публичной части
package restapi

import (
	"net/http"
	"visit/src/model"
	"visit/src/service"
	"visit/src/util/helper"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Tags Пользователь
// @Summary Регистрация нового пользователя
// @Description Регистрация нового пользователя в системе медицинской организации
// @Accept json
// @Produce json
// @Param request body model.UserRegisterRequest true "Данные для регистрации"
// @Success 201 {object} model.TokenResponse "Токен авторизации"
// @Failure 400 {object} helper.ErrorResponse "Ошибка валидации"
// @Failure 409 {object} helper.ErrorResponse "Пользователь уже существует"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/register [post]
func Register(ctx *gin.Context) {
	var req model.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Message:      "VALIDATION_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	user, err := service.RegisterUser(req)
	if err != nil {
		ctx.JSON(http.StatusConflict, helper.ErrorResponse{
			Message:      "REGISTER_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	token, err := helper.GenerateJWT(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Message:      "TOKEN_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.TokenResponse{Token: token})
}

// Login godoc
// @Tags Пользователь
// @Summary Авторизация пользователя
// @Description Авторизация пользователя по email и паролю
// @Accept json
// @Produce json
// @Param request body model.UserLoginRequest true "Данные для авторизации"
// @Success 200 {object} model.TokenResponse "Токен авторизации"
// @Failure 400 {object} helper.ErrorResponse "Ошибка валидации"
// @Failure 401 {object} helper.ErrorResponse "Неверные учётные данные"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/login [post]
func Login(ctx *gin.Context) {
	var req model.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Message:      "VALIDATION_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	user, err := service.LoginUser(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, helper.ErrorResponse{
			Message:      "AUTH_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	token, err := helper.GenerateJWT(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Message:      "TOKEN_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.TokenResponse{Token: token})
}

// GetProfile godoc
// @Tags Личный кабинет
// @Summary Получение профиля пользователя
// @Description Получение данных профиля текущего авторизованного пользователя
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.UserProfileResponse "Профиль пользователя"
// @Failure 401 {object} helper.ErrorResponse "Не авторизован"
// @Failure 404 {object} helper.ErrorResponse "Пользователь не найден"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/profile [get]
func GetProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("userUuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, helper.ErrorResponse{
			Message:      "AUTH_ERROR",
			ErrorMessage: "пользователь не авторизован",
		})
		return
	}

	profile, err := service.GetUserProfile(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.ErrorResponse{
			Message:      "NOT_FOUND",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// UpdateProfile godoc
// @Tags Личный кабинет
// @Summary Обновление профиля пользователя
// @Description Обновление данных профиля текущего авторизованного пользователя
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.UserUpdateRequest true "Данные для обновления"
// @Success 200 {object} model.UserProfileResponse "Обновлённый профиль"
// @Failure 400 {object} helper.ErrorResponse "Ошибка валидации"
// @Failure 401 {object} helper.ErrorResponse "Не авторизован"
// @Failure 500 {object} helper.ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/profile [put]
func UpdateProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("userUuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, helper.ErrorResponse{
			Message:      "AUTH_ERROR",
			ErrorMessage: "пользователь не авторизован",
		})
		return
	}

	var req model.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse{
			Message:      "VALIDATION_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	profile, err := service.UpdateUserProfile(userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{
			Message:      "UPDATE_ERROR",
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
