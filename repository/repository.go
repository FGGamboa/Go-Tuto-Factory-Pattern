package repository

import (
	"fmt"
	"tuto-factory-pattern/configuration"
	"tuto-factory-pattern/repository/mysql"
	"tuto-factory-pattern/repository/sqlserver"
)

type Repository interface {
	Find(id int) string
	Save(data string) error
}

func New(config *configuration.Configuration) (Repository, error) {

	var repo Repository
	var err error

	switch config.Engine {
	case "mysql":
		repo = mysql.New()
	case "sqlserver":
		repo = sqlserver.New()
	default:
		err = fmt.Errorf("invalid engine: %s", config.Engine)

	}
	return repo, err
}
