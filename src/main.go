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

	todoHandler := handlers.NewTodo()
	bookHandler := handlers.NewBook()

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.GET("/", todoHandler.GetTodos)
	router.GET("/books", bookHandler.GetBooks)

	router.POST("/", todoHandler.CreateTodo)

	router.PUT("/:id", todoHandler.UpdateTodo)

	router.DELETE("/:id", todoHandler.DeleteTodo)

	router.Run(":9090")

}
