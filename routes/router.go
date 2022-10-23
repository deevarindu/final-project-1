package routes

import (
	"github.com/deevarindu/final-project-1/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.HealthCheck)

	return router
}
