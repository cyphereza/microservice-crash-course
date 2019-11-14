package repository

import "github.com/cyphereza/microservice-crash-course/model"

type User interface {
	Create(name string) (string, error)
	RetrieveUserByName(name string) ([]model.User, error)
	RetrieveUserByID(id string) (string, error)
}
