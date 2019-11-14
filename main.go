package main

import (
	"database/sql"
	"net/http"

	userHandler "github.com/cyphereza/microservice-crash-course/handler/user"
	userRepo "github.com/cyphereza/microservice-crash-course/repository/user"
	userUsecase "github.com/cyphereza/microservice-crash-course/usecase/user"
	"github.com/labstack/echo"
)

var mdb *sql.DB

func main() {
	e := echo.New()

	// Tell echo to use these handlers, usecases and repos
	repoUser := userRepo.NewMysqlRepository(mdb)
	usecaseUser := userUsecase.NewUserUsecase(repoUser)
	userHandler.NewUserHandler(e, usecaseUser)

	s := &http.Server{
		Addr: ":7000",
	}

	e.Logger.Fatal(e.StartServer(s))
}
