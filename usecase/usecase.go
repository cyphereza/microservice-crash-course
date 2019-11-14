package usecase

type User interface {
	CreateUser(name string) (string, error)
	RetrieveUser(name string) (string, error)
}
