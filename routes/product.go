package routes

import (
	"backend/controllers"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.ConnDB)
	h := controllers.ControlProduct(productRepository)

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProducts)
	e.POST("/product", h.CreateProduct)
	e.PATCH("/product/:id", h.UpdateProduct)
	e.DELETE("/product/:id", h.DeleteProduct)
}