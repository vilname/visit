package config

import (
	"visit/docs"
	restadmin "visit/src/controller/restAdmin"
	restapi "visit/src/controller/restApi"
	"visit/src/util/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRoute роутинг
func InitRoute() *gin.Engine {
	router := gin.New()
	//router.Use(middleware.EnableCORS)
	router.Use(middleware.AuthenticationMiddleware)

	admin := router.Group(`/admin`)
	admin.GET("/doctor", restadmin.Doctor)

	api := router.Group(`/api`)
	api.GET("/doctor", restapi.Doctor)

	// Пользователь: регистрация и авторизация (без токена)
	api.POST("/user/register", restapi.Register)
	api.POST("/user/login", restapi.Login)

	// Личный кабинет пользователя (требуется токен)
	api.GET("/user/profile", restapi.GetProfile)
	api.PUT("/user/profile", restapi.UpdateProfile)

	// Врачи
	api.GET("/doctors", restapi.GetDoctors)
	api.GET("/doctors/:id", restapi.GetDoctorByID)

	// Записи на приём (требуется токен)
	api.POST("/user/appointments", restapi.CreateAppointment)
	api.GET("/user/appointments", restapi.GetAppointments)
	api.DELETE("/user/appointments/:id", restapi.CancelAppointment)

	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Title = "Visit - API медицинской организации"
	docs.SwaggerInfo.Description = "Сервис для записи на приём к врачам медицинской организации"
	docs.SwaggerInfo.Version = "1.0"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
