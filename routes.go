package main

import (
	"github.com/gin-gonic/gin"
	"todo-backend/handlers"
	"github.com/gin-contrib/cors"
	"os"
)

func configureCors(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowMethods("DELETE", "PATCH")
	router.Use(cors.New(corsConfig))
}

func setupRoutes(todos *handlers.Todos) *gin.Engine {
	router := gin.Default()

	configureCors(router)
	gin.DisableConsoleColor()

	router.GET("/:id", readTodo(todos))
	router.GET("/", listTodos(todos))

	router.POST("/", addTodo(todos))
	router.PATCH("/:id", updateTodo(todos))
	router.DELETE("/", deleteAllTodos(todos))
	router.DELETE("/:id", deleteTodo(todos))

	return router
}

func listTodos(todos *handlers.Todos) func(context *gin.Context) {
	return func(context *gin.Context) {
		allTodos := todos.List()
		context.JSON(200, allTodos)
	}
}

func readTodo(todos *handlers.Todos) func(context *gin.Context) {
	return func(context *gin.Context) {
		if t := todos.Find(context.Param("id")); t != nil {
			context.JSON(200, t)
		} else {
			context.Status(404)
		}
	}
}

func addTodo(todos *handlers.Todos) func(context *gin.Context) {
	return func(context *gin.Context) {
		var todo handlers.Todo
		if e := context.BindJSON(&todo); e == nil {
			newTodo := todos.Add(todo, basePath())
			context.JSON(201, newTodo)
		} else {
			context.AbortWithError(400, e)
		}
	}
}

func updateTodo(todos *handlers.Todos) func(context *gin.Context) {
	return func(context *gin.Context) {
		var todo handlers.Todo
		if t := todos.Find(context.Param("id")); t != nil {
			if e := context.BindJSON(&todo); e == nil {
				t.Update(todo)
				context.JSON(200, t)
			} else {
				context.AbortWithError(400, e)
			}
		} else {
			context.Status(404)
		}
	}
}

func deleteAllTodos(todos *handlers.Todos) func(context *gin.Context) {
	return func(context *gin.Context) {
		todos.RemoveAll()
		context.Status(200)
	}
}

func deleteTodo(todos *handlers.Todos) func(context *gin.Context) {
	return func(context *gin.Context) {
		todos.RemoveAll()
		context.Status(200)
	}
}

func basePath() string {
	baseUrl := os.Getenv("BASE_URL")
	if len(baseUrl) == 0 {
		return "http://localhost:8080"
	}
	return baseUrl
}
