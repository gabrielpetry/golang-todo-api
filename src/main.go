package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"host.local/go/golang-todo-api/src/database"
	"host.local/go/golang-todo-api/src/handlers"
	"host.local/go/golang-todo-api/src/middlewares"
)

func main() {
	log.Info("Starting the application")
	database.Init()

	todoController := handlers.NewTodo()

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.GET("/", todoController.GetTodos)

	router.POST("/", todoController.CreateTodo)

	router.PUT("/:id", todoController.UpdateTodo)

	router.DELETE("/:id", todoController.DeleteTodo)

	router.Run(":9090")

}
