package usecase

import "github.com/cyphereza/microservice-crash-course/model"

type User interface {
	CreateUser(name string) (string, error)
	RetrieveUser(name string) ([]model.User, error)
}
