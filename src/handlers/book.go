package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"host.local/go/golang-todo-api/src/models"
)

type Books struct{}

var bookModel models.Todo = models.Todo{}

// NewTodo creates an instance of Books
func NewBook() *Books {
	return &Books{}
}

func (p *Books) CreateBook(c *gin.Context) {

	c.BindJSON(&bookModel)

	if len(bookModel.Task) < 2 {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "todo must be at least 3 chars length"},
		)
		return
	}

	err := bookModel.InsertOne()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusCreated)
}

func (todo *Books) GetBooks(c *gin.Context) {
	bookModel, err := bookModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"tasks": bookModel})
	c.JSON(http.StatusOK, bookModel)
}

func (todo *Books) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	c.BindJSON(&bookModel)

	err := bookModel.Update(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusNoContent)
}

func (todo *Books) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := bookModel.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return

	}

	c.Status(http.StatusNoContent)
}
