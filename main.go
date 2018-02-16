package main

import (
	"github.com/gin-gonic/gin"
	"todo-backend/handlers"
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

func main() {
	SetupRoutes().Run(":8080")
}


