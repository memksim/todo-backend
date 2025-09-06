package http

import (
	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"status": "ok"})
}
