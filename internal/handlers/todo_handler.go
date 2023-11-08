package handlers

import (
	"net/http"

	"github.com/doganarif/todo/internal/models"
	"github.com/doganarif/todo/internal/repositories"
	"github.com/gin-gonic/gin"
)

type ITodoHandler interface {
	CreateTodo() gin.HandlerFunc
}

type TodoHandler struct {
	todoRepository repositories.ITodoRepository
}

func NewTodoHandler(repo repositories.ITodoRepository) ITodoHandler {
	return &TodoHandler{todoRepository: repo}

}

func (t TodoHandler) CreateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo
		c.BindJSON(&todo)

		err := t.todoRepository.Create(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"s": "err"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"s": "ok"})
		return
	}
}
