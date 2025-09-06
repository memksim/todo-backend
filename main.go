package main

import (
	"github.com/memksim/todo-backend/internal/http"
)

func main() {
	err := http.SetupRouter()
	if err != nil {
		panic(err)
	}
}
