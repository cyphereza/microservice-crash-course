package user

import (
	"database/sql"
	"log"

	"github.com/cyphereza/microservice-crash-course/model"
	"github.com/cyphereza/microservice-crash-course/repository"
	_ "github.com/go-sql-driver/mysql"
)

type mysqlRepository struct {
	sess *sql.DB
}

func NewMysqlRepository(mdb *sql.DB) repository.User {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/test")
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &mysqlRepository{
		sess: db,
	}
}

func (mdb *mysqlRepository) Create(name string) (string, error) {
	return "", nil
}

func (mdb *mysqlRepository) RetrieveUserByName(name string) (output []model.User, err error) {
	results, err := mdb.sess.Query("SELECT id, name FROM user WHERE name like '%" + name + "%'")
	if err != nil {
		log.Fatalln(err.Error())
	}

	for results.Next() {
		var idOutput int
		var nameString string

		err = results.Scan(&idOutput, &nameString)
		if err != nil {
			log.Fatalln(err.Error())
		}

		row := new(model.User)
		row.ID = idOutput
		row.Name = nameString
		output = append(output, *row)
	}

	return output, nil
}

func (mdb *mysqlRepository) RetrieveUserByID(id string) (string, error) {
	return "", nil
}
