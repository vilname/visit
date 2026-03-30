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

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
