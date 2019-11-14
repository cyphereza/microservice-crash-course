package handler

import "github.com/labstack/echo"

type User interface {
	createUser(c echo.Context) error
	retrieveUser(c echo.Context) error
}
