package main

import "todo-backend/handlers"

func main() {
	todos := handlers.Todos{}
	setupRoutes(&todos).Run()
}


