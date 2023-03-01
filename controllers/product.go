package controllers

import (
	productsdto "backend/dto/product"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type productControl struct {
	ProductRepository repositories.ProductRepository
}

func ControlProduct(ProductRepository repositories.ProductRepository) *productControl {
	return &productControl{ProductRepository}
}

func (h *productControl) FindProducts(c echo.Context) error {
	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ProductResult{Status: http.StatusOK, Data: products})
}

func (h *productControl) GetProducts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	products, err := h.ProductRepository.GetProducts(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ProductResult{Status: http.StatusOK, Data: convProduct(products)})
}

func (h *productControl) CreateProduct(c echo.Context) error {
	request := new(productsdto.CreateProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// data form pattern submit to pattern entity db product
	product := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Stock:       request.Stock,
		Photo:       request.Photo,
	}

	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ProductResult{Status: http.StatusOK, Data: convProduct(data)})
}

func (h *productControl) UpdateProduct(c echo.Context) error {
	request := new(productsdto.UpdateProductRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProducts(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Price != 0 {
		product.Price = request.Price
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.Stock != 0 {
		product.Stock = request.Stock
	}
	if request.Photo != "" {
		product.Photo = request.Photo
	}

	data, err := h.ProductRepository.UpdateProduct(product, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ProductResult{Status: http.StatusOK, Data: convProduct(data)})
}

func (h *productControl) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.ProductRepository.GetProducts(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ProductRepository.DeleteProduct(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ProductResult{Status: http.StatusOK, Data: convProduct(data)})
}

func convProduct(u models.Product) productsdto.ProductResponse {
	return productsdto.ProductResponse{
		ID:          u.ID,
		Name:        u.Name,
		Price:       u.Price,
		Description: u.Description,
		Stock:       u.Stock,
		Photo:       u.Photo,
	}
}
