package http

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func SetupRouter() error {
	router = gin.Default()

	setupBackendRouter(router)
	setupTaskRouter(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return err
	}
	return nil
}

func setupBackendRouter(r *gin.Engine) {
	r.GET("/health", GetHealth)
}

func setupTaskRouter(r *gin.Engine) {
	r.GET("/tasks", GetTasks)
	r.POST("/task", PostTask)
}
