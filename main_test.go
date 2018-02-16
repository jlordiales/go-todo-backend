package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func Test_rootPathShouldRespondWith200(t *testing.T) {
	testRouter := routes()
	req, _ := http.NewRequest("GET", "/", nil)

	resp := performRequest(testRouter, req)

	assert.Equal(t, 200, resp.Code)
}

func Test_rootPathShouldAcceptPostWithNewTodo(t *testing.T) {
	testRouter := routes()
	json := `{"title": "a todo"}`
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(json)))

	resp := performRequest(testRouter, req)

	assert.Equal(t, 200, resp.Code)
}

func Test_getAfterPostReturnsAllTodos(t *testing.T) {
	testRouter := routes()
	json := `{"title": "a todo"}`
	deleteAll(testRouter)
	addTodo(json, testRouter)

	listRequest, _ := http.NewRequest("GET", "/", nil)
	resp := performRequest(testRouter, listRequest)

	assert.Equal(t, 200, resp.Code)
}

func Test_deleteShouldRemoveAllTodos(t *testing.T) {
	testRouter := routes()
	json := `{"title": "a todo"}`
	addTodo(json, testRouter)

	resp := deleteAll(testRouter)

	assert.Equal(t, 200, resp.Code)

	listRequest, _ := http.NewRequest("GET", "/", nil)
	listResponse := performRequest(testRouter, listRequest)
	assert.JSONEq(t, "[]", listResponse.Body.String())
}

func Test_newTodosShouldBeCreatedAsNotCompleted(t *testing.T) {
	testRouter := routes()
	json := `{"title": "a todo"}`
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(json)))

	resp := performRequest(testRouter, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(),`"completed":false`)
}

func Test_newTodosShouldHaveAnUrl(t *testing.T) {
	testRouter := routes()
	json := `{"title": "a todo"}`
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(json)))

	resp := performRequest(testRouter, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(),`"url":`)
}

func deleteAll(testRouter *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("DELETE", "/", nil)
	return performRequest(testRouter, req)
}

func routes() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return SetupRoutes()
}

func addTodo(json string, testRouter *gin.Engine)  {
	createTodoReq, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(json)))
	performRequest(testRouter, createTodoReq)
}

func performRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}