package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	setupBackendRouter(r)
	setupTaskRouter(r)
	return r
}

func setupBackendRouter(r *gin.Engine) {
	r.GET("/health", GetHealth)
}

func setupTaskRouter(r *gin.Engine) {
	r.GET("/tasks", GetTasks)
	r.POST("/task", PostTask)
}
