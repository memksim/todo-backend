package api

import (
	"github.com/gin-gonic/gin"
	"github.com/memksim/todo-backend/internal/api/dto"
	"net/http"
)

func GetTasks(c *gin.Context) {
	tasks := []dto.TaskResponse{
		{0, "Закончить проект", "2025-10-01T18:36:21+05:00", false},
		{1, "Убраться дома", "2025-09-04T09:30:00+05:00", true},
		{2, "Заказать справку", "2025-09-5T12:12:12+05:00", false},
	}
	c.JSON(http.StatusOK, tasks)
}

// PostTask example:
// curl -X POST http://localhost:8080/task -H 'Content-Type: application/json' -d '{"title":"Send email", "expires_at":"2025-09-16T21:45:21+05:00"}'
func PostTask(c *gin.Context) {
	var newTaskInfo dto.CreateTaskRequest
	if err := c.BindJSON(&newTaskInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask := dto.TaskResponse{
		Id:         11,
		Title:      newTaskInfo.Title,
		ExpiresAt:  newTaskInfo.ExpiresAt,
		IsComplete: false,
	}
	c.JSON(http.StatusCreated, newTask)
}
