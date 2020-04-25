package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"host.local/go/golang-todo-api/src/database"
	"host.local/go/golang-todo-api/src/handlers"
)

func main() {
	log.Info("Starting the application")
	database.Init()

	todoController := handlers.NewTodo()

	router := gin.Default()

	router.GET("/", todoController.GetTodos)

	router.POST("/", todoController.CreateTodo)

	router.PUT("/:id", todoController.UpdateTodo)

	router.DELETE("/:id", todoController.DeleteTodo)

	router.Run(":9090")

}
