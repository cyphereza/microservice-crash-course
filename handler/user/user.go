package user

import (
	"net/http"

	"github.com/cyphereza/microservice-crash-course/model"
	"github.com/cyphereza/microservice-crash-course/usecase"
	"github.com/labstack/echo"
)

type userHandler struct {
	usecase usecase.User
}

func NewUserHandler(e *echo.Echo, usecase usecase.User) {
	handler := &userHandler{
		usecase: usecase,
	}

	e.POST("/user/create", handler.createUser)
	e.GET("/user/retrieve/:name1", handler.retrieveUser)
}

func (h *userHandler) createUser(c echo.Context) error {
	return nil
}

func (h *userHandler) retrieveUser(c echo.Context) error {
	name := c.Param("name1")
	result, err := h.usecase.RetrieveUser(name)
	if err != nil || len(result) < 1 {
		returnMessage := new(model.ErrorReturn)
		returnMessage.ErrorMessage = "Error"
		return c.JSON(http.StatusOK, returnMessage)
	}

	return c.JSON(http.StatusOK, result)
}
