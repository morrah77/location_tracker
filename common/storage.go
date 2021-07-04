package common

type Storage interface {
	Fetch(string, int) (interface{}, error)
	Store(string, interface{}) (error)
	Delete(string) (error)
}
