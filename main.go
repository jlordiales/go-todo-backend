package main

import "todo-backend/handlers"

func main() {
	todos := handlers.Todos{}
	routes := setupRoutes(&todos)
	routes.Run()
}


