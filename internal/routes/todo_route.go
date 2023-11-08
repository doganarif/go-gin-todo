package routes

import (
	"github.com/doganarif/todo/internal/handlers"
	"github.com/doganarif/todo/internal/repositories"
	"github.com/gin-gonic/gin"
)

func TodoGroup(router *gin.RouterGroup, todoRepo repositories.ITodoRepository) {
	todoHandler := handlers.NewTodoHandler(todoRepo)

	todoRoutes := router.Group("/todo")
	{
		todoRoutes.POST("/", todoHandler.CreateTodo())
	}
}
