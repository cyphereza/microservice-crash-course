package user

import (
	"github.com/cyphereza/microservice-crash-course/model"
	"github.com/cyphereza/microservice-crash-course/repository"
	"github.com/cyphereza/microservice-crash-course/usecase"
)

type userUsecase struct {
	repo repository.User
}

func NewUserUsecase(repo repository.User) usecase.User {
	return &userUsecase{
		repo: repo,
	}
}

func (uc *userUsecase) CreateUser(name string) (string, error) {
	return "", nil
}

func (uc *userUsecase) RetrieveUser(name string) ([]model.User, error) {
	result, _ := uc.repo.RetrieveUserByName(name)
	return result, nil
}
