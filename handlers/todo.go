package handlers

import "github.com/gin-gonic/gin"


var todos = make([]Todo, 0)

type Todo struct {
	Title string `json:"title" `
}

func ListTodos(c *gin.Context) {
	c.JSON(200, todos)
}

func DeleteTodos(c *gin.Context)  {
	todos = make([]Todo, 0)
	c.Status(200)
}

func AddTodo(c *gin.Context) {
	var todo Todo
	if e := c.BindJSON(&todo); e == nil {
		todos = append(todos, todo)
		c.JSON(200, todo)
	} else {
		c.AbortWithError(400, e)
	}
}
