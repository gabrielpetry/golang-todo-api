package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"host.local/go/golang-todo-api/src/models"
)

type Todos struct{}

var todoModel models.Todo = models.Todo{}

// NewTodo creates an instance of todos
func NewTodo() *Todos {
	return &Todos{}
}

func (p *Todos) CreateTodo(c *gin.Context) {

	c.BindJSON(&todoModel)

	if len(todoModel.Task) < 2 {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "todo must be at least 3 chars length"},
		)
		return
	}

	err := todoModel.InsertOne()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusCreated)
}

func (todo *Todos) GetTodos(c *gin.Context) {
	todoModel, err := todoModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"tasks": todoModel})
	c.JSON(http.StatusOK, todoModel)
}

func (todo *Todos) UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	c.BindJSON(&todoModel)

	err := todoModel.Update(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusNoContent)
}

func (todo *Todos) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	err := todoModel.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return

	}

	c.Status(http.StatusNoContent)
}
