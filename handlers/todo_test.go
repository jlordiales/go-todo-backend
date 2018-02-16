package handlers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTodos_FindReturnsNilForNonExistingId(t *testing.T) {
	todos := Todos{}

	assert.Nil(t, todos.Find("id"))
}

func TestTodos_FindReturnsCorrectTodo(t *testing.T) {
	todos := Todos{}
	addedTodo := todos.Add(Todo{Title: "awesome title"}, "localhost")

	assert.Equal(t, "awesome title", todos.Find(addedTodo.Id).Title)
}

func TestTodos_ListReturnsAllTodos(t *testing.T) {
	todos := Todos{}
	todos.Add(Todo{Title: "todo 1"}, "localhost")
	todos.Add(Todo{Title: "todo 2"}, "localhost")

	assert.Len(t, todos.List(), 2)
}

func TestTodos_RemoveAll(t *testing.T) {
	todos := Todos{}
	todos.Add(Todo{Title: "todo 1"}, "localhost")
	todos.Add(Todo{Title: "todo 2"}, "localhost")

	assert.Len(t, todos.List(), 2)

	todos.RemoveAll()
	assert.Len(t, todos.List(), 0)
}

