package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"todo-backend/handlers"
	"encoding/json"
)

func Test_rootPathShouldRespondWith200(t *testing.T) {
	testRouter := routes()
	req, _ := http.NewRequest("GET", "/", nil)

	resp := performRequest(testRouter, req)

	assert.Equal(t, 200, resp.Code)
}

func Test_rootPathShouldAcceptPostWithNewTodo(t *testing.T) {
	testRouter := routes()

	resp, createdTodo := createTodo(`{"title": "a todo"}`, testRouter)

	assert.Equal(t, 201, resp.Code)
	assert.NotEmpty(t, createdTodo.Url)
	assert.Equal(t, "a todo", createdTodo.Title)
	assert.False(t, createdTodo.Completed)
}

func Test_getAfterPostReturnsAllTodos(t *testing.T) {
	testRouter := routes()
	createTodo(`{"title": "a todo"}`, testRouter)
	createTodo(`{"title": "another todo"}`, testRouter)

	resp, todos := getAllTodos(testRouter)

	assert.Equal(t, 200, resp.Code)
	assert.Len(t, todos, 2)
}

func Test_deleteShouldRemoveAllTodos(t *testing.T) {
	testRouter := routes()
	createTodo(`{"title": "a todo"}`, testRouter)

	resp := deleteAll(testRouter)
	assert.Equal(t, 200, resp.Code)

	_, todos := getAllTodos(testRouter)
	assert.Len(t, todos, 0)
}

func Test_shouldBePossibleToGetATodoByItsUrl(t *testing.T) {
	testRouter := routes()
	_, createdTodo := createTodo(`{"title": "a todo"}`, testRouter)

	resp, todo := getTodo(testRouter, "/" + createdTodo.Id)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "a todo", todo.Title)
}

func Test_canUpdateTodoByPatchingToItsUrl(t *testing.T) {
	testRouter := routes()
	_, createdTodo := createTodo(`{"title": "a todo"}`, testRouter)

	resp, todo := patchTodo(`{"title": "a new title", "completed": true}`, createdTodo.Id, testRouter)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "a new title", todo.Title)
	assert.True(t, todo.Completed)
}

func deleteAll(testRouter *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("DELETE", "/", nil)
	return performRequest(testRouter, req)
}

func routes() *gin.Engine {
	gin.SetMode(gin.TestMode)
	todos := handlers.Todos{}
	return setupRoutes(&todos)
}

func performRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func createTodo(todoJson string, engine *gin.Engine) (*httptest.ResponseRecorder, handlers.Todo) {
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(todoJson)))
	resp := performRequest(engine, req)

	var createdTodo handlers.Todo
	json.Unmarshal(resp.Body.Bytes(), &createdTodo)
	return resp, createdTodo
}

func patchTodo(todoJson string, id string, engine *gin.Engine) (*httptest.ResponseRecorder, handlers.Todo) {
	req, _ := http.NewRequest("PATCH", "/" + id, bytes.NewBuffer([]byte(todoJson)))
	resp := performRequest(engine, req)

	var createdTodo handlers.Todo
	json.Unmarshal(resp.Body.Bytes(), &createdTodo)
	return resp, createdTodo
}

func getAllTodos(engine *gin.Engine) (*httptest.ResponseRecorder, []handlers.Todo) {
	listRequest, _ := http.NewRequest("GET", "/", nil)
	resp := performRequest(engine, listRequest)

	var allTodos []handlers.Todo
	json.Unmarshal(resp.Body.Bytes(), &allTodos)
	return resp, allTodos
}

func getTodo(engine *gin.Engine, id string) (*httptest.ResponseRecorder, handlers.Todo) {
	listRequest, _ := http.NewRequest("GET", id, nil)
	resp := performRequest(engine, listRequest)

	var todo handlers.Todo
	json.Unmarshal(resp.Body.Bytes(), &todo)
	return resp, todo
}
