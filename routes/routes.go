package routes

import "github.com/labstack/echo/v4"

func Routes(e *echo.Group) {
	TodoRoutes(e)
	UserRoutes(e)
	ProductRoutes(e)
}