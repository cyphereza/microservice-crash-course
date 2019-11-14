package repository

type User interface {
	Create(name string) (string, error)
	RetrieveUserByName(name string) (string, error)
	RetrieveUserByID(id string) (string, error)
}
