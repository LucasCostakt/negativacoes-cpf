package repository

type Storage interface {
	CreatePrincipalTable() error
}
