package repository

import (
	"service-import-data/service"

	_ "github.com/go-sql-driver/mysql"
)


const (
	Username = "root"
	Password = "admin"
	Hostname = "mysql:3306"
	DBName   = "testeDB"
)

type Storage interface {
	SaveJson(data []*service.Data) error
}