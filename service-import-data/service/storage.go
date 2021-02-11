package service

type Storage interface {
	ConvertFileToStruct(namePath string) ([]*Data, error)
}