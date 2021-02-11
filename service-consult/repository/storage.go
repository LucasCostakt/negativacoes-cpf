package repository

import (
	"service-consult/service"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Username = "root"
	Password = "admin"
	Hostname = "mysql:3306"
	DBName   = "testeDB"
)

type Storage interface {
	ConsultNegativacoes(data string) ([]*service.Data, error)
}
