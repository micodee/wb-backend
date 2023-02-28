package controllers

import (
	"backend/entities"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// menampilkan semua data todo
func FindTodos(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(200)
	return json.NewEncoder(c.Response()).Encode(entities.Todos)
}

// menampilkan berdasarkan id
func GetTodo(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// menangkap url
	id := c.Param("id")

	var todoData entities.TodosStruct
	var isGetTodo = false

	// mengecek apakah idnya ada
	for _, todo := range entities.Todos {
		if id == todo.Id {
			isGetTodo = true
			todoData = todo
		}
	}

	if !isGetTodo {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID : " + id + " not found")
	}

	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(todoData)
}

// menambahkan data todo
func CreateTodo(c echo.Context) error {
	var data entities.TodosStruct

	json.NewDecoder(c.Request().Body).Decode(&data)

	entities.Todos = append(entities.Todos, data)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(entities.Todos)
}

func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	var data entities.TodosStruct
	var isGetTodo = false

	json.NewDecoder(c.Request().Body).Decode(&data)

	for idx, todo := range entities.Todos {
		if id == todo.Id {
			isGetTodo = true
			entities.Todos[idx] = data
		}
	}

	if !isGetTodo {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID : " + id + " not found")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(entities.Todos)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	var isGetTodo = false
	var index = 0

	for idx, todo := range entities.Todos {
		if id == todo.Id {
			isGetTodo = true
			index = idx
		}
	}

		if !isGetTodo {
			c.Response().WriteHeader(http.StatusNotFound)
			return json.NewEncoder(c.Response()).Encode("ID : " + id + " not found")
		}

		entities.Todos = append(entities.Todos[:index], entities.Todos[index+1:]...)
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(entities.Todos)
}