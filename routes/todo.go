package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func TodoRoutes(e *echo.Group) {
	e.GET("/todos", controllers.FindTodos)
	e.GET("/todos/:id", controllers.GetTodo)
	e.POST("/todos", controllers.CreateTodo)
	e.PATCH("/todos/:id", controllers.UpdateTodo)
	e.DELETE("/todos/:id", controllers.DeleteTodo)
}