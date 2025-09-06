package http

import (
	"github.com/gin-gonic/gin"
	"github.com/memksim/todo-backend/internal/http/dto"
	"net/http"
)

func GetTasks(c *gin.Context) {
	tasks := []dto.Task{
		{0, "Закончить проект", false, 1758006000000},
		{1, "Убраться дома", true, 1756969200000},
		{2, "Заказать справку", false, 1757154600000},
	}
	c.JSON(http.StatusOK, tasks)
}

// PostTask example:
// curl -X POST http://localhost:8080/task -H 'Content-Type: application/json' -d '{"id":23, "title":"Send email", "is_completed":false, "expires_at":1759746600000}'
func PostTask(c *gin.Context) {
	var newTask dto.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, newTask)
}
