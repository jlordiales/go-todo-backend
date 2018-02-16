package handlers

import (
	//"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	//"errors"
)

type Todo struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Url       string `json:"url,omitempty"`
	Order     int    `json:"order,omitempty"`
}

type Todos map[string]*Todo

func (t Todos) List() []*Todo {
	var todos = make([]*Todo, 0)
	for _, v := range t {
		todos = append(todos, v)
	}
	return todos
}

func (todos Todos) Find(id string) *Todo {
	t, ok := todos[id]
	if ok {
		return t
	}
	return nil
}

func (t Todos) Add(todo Todo, basePath string) *Todo {
	todo.Completed = false
	todo.Id = uuid.Must(uuid.NewV4()).String()
	todo.Url = basePath + "/" + todo.Id
	t[todo.Id] = &todo
	return &todo
}

func (t Todos) RemoveAll() {
	for k := range t {
		delete(t, k)
	}
}

func (t Todos) Remove(id string) {
	delete(t, id)
}

func (todo *Todo) Update(newTodo Todo) *Todo {
	todo.Title = newTodo.Title
	todo.Completed = newTodo.Completed
	todo.Order = newTodo.Order
	return todo
}
