package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

type Todos struct{}

var todoModel *models.Todo = models.NewTodo()

// NewTodo creates an instance of todos
func NewTodo() *Todos {
	return &Todos{}
}
func (p *Todos) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	newTodo := models.Todo{
		Task: "teste",
	}
	// var task models.Todo
	//_ := json.NewDecoder(r.Body)
	// fmt.Println(newTodo)
	newTodo.InsertOne()
	json.NewEncoder(rw).Encode(newTodo)
}

func (p *Todos) GetTodos(rw http.ResponseWriter, r *http.Request) {
	lp, _ := todoModel.GetAll()
	encoder := json.NewEncoder(rw)
	fmt.Println(lp)
	encoder.Encode(lp)
	// err := lp.ToJSON(rw)
	// if err != nil {
	// 	http.Error(rw, "Unabble to marshal json", http.StatusInternalServerError)
	// }
}
