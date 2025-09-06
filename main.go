package main

import "github.com/memksim/todo-backend/internal/api"

func main() {
	router := api.SetupRouter()

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
