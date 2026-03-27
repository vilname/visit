package config

import (
	"visit/docs"
	"visit/src/controller/restAdmin"
	"visit/src/controller/restApi"
	"visit/src/util/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	//router.Use(middleware.EnableCORS)
	router.Use(middleware.AuthenticationMiddleware)

	admin := router.Group(`/admin`)
	admin.GET("/doctor", restAdmin.Doctor)

	api := router.Group(`/api`)
	api.GET("/doctor", restApi.Doctor)

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
