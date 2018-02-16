package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/location"
	"github.com/satori/go.uuid"
	"errors"
)

var todos = make([]Todo, 0)

type Todo struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Url       string `json:"url,omitempty"`
}

func ListTodos(c *gin.Context) {
	c.JSON(200, todos)
}

func ReadTodo(c *gin.Context) {
	id := c.Param("id")
	if todo, e := findTodo(id); e == nil {
		c.JSON(200, todo)
	} else {
		c.Status(404)
	}
}

func findTodo(id string) (*Todo, error) {
	for _, t := range todos {
		if t.Id == id {
			return &t, nil
		}
	}
	return nil, errors.New("could not find todo")
}

func DeleteTodos(c *gin.Context) {
	todos = make([]Todo, 0)
	c.Status(200)
}

func AddTodo(c *gin.Context) {
	var todo Todo
	if e := c.BindJSON(&todo); e == nil {
		todo.Completed = false
		todo.Id = uuid.Must(uuid.NewV4()).String()
		todo.Url = location.Get(c).String() + "/" + todo.Id

		todos = append(todos, todo)
		c.JSON(200, todo)
	} else {
		c.AbortWithError(400, e)
	}
}
