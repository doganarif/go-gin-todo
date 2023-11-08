package main

import (
	"github.com/doganarif/todo/internal/repositories"
	"github.com/doganarif/todo/internal/routes"
	"github.com/doganarif/todo/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.Initialize()

	apiRoute := router.Group("/api/v1/")

	todoRepo := repositories.NewTodoRepository(db)

	routes.TodoGroup(apiRoute, todoRepo)

	//
	router.Run()
}
