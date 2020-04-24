package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"host.local/go/golang-todo-api/src/database"
	"host.local/go/golang-todo-api/src/handlers"
)

func main() {
	log.Println("Starting the application")
	database.Init()

	todoController := handlers.NewTodo()

	router := gin.Default()

	router.GET("/", todoController.GetTodos)

	router.POST("/", todoController.CreateTodo)

	router.PUT("/:id", todoController.UpdateTodo)

	router.DELETE("/:id", todoController.DeleteTodo)

	router.Run(":9090")

	// ph := handlers.NewTodo()
	// sm := mux.NewRouter()

	// getRouter := sm.Methods("GET").Subrouter()
	// getRouter.HandleFunc("/", ph.GetTodos)

	// postRouter := sm.Methods("POST").Subrouter()
	// postRouter.HandleFunc("/", ph.CreateTodo)
	// // postRouter.Use(ph.MiddlewareTodoValidation)

	// putRouter := sm.Methods("PUT").Subrouter()
	// putRouter.HandleFunc("/{id}", ph.UpdateTodo)
	// putRouter.Use(ph.MiddlewareTodoValidation)

	// deleteRouter := sm.Methods("DELETE").Subrouter()
	// deleteRouter.HandleFunc("/{id}", ph.DeleteTodo)

	// corsOpts := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
	// 	AllowedMethods: []string{
	// 		http.MethodGet, //http methods for your app
	// 		http.MethodPost,
	// 		http.MethodPut,
	// 		http.MethodPatch,
	// 		http.MethodDelete,
	// 		http.MethodOptions,
	// 		http.MethodHead,
	// 	},
	// })
	// server := &http.Server{
	// 	Addr:         ":9090",
	// 	Handler:      corsOpts.Handler(sm),
	// 	IdleTimeout:  120 * time.Second,
	// 	ReadTimeout:  1 * time.Second,
	// 	WriteTimeout: 1 * time.Second,
	// }

	// go func() {
	// 	err := server.ListenAndServe()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		os.Exit(1)
	// 	}
	// }()

	// // trap sigterm or interupt and gracefully shutdown the server
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// signal.Notify(c, os.Kill)

	// // block until a signal is received.
	// sig := <-c
	// log.Println("got signal:", sig)

	// // gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// server.Shutdown(ctx)
}
