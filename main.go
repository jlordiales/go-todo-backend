package main

import (
	"github.com/gin-gonic/gin"
	"todo-backend/handlers"
	"os"
)


func SetupRoutes() *gin.Engine {
	router := gin.Default()
	gin.DisableConsoleColor()

	router.GET("/ping", handlers.Ping)
	router.GET("/", handlers.ListTodos)

	router.POST("/", handlers.AddTodo)
	router.DELETE("/", handlers.DeleteTodos)

	return router
}

func port() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8080"
	}
	return ":" + port
}

func main() {
	SetupRoutes().Run(port())
}


