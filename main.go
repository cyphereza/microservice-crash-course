package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	s := &http.Server{
		Addr: ":7000",
	}

	e.Logger.Fatal(e.StartServer(s))
}
