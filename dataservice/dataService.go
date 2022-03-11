package dataservice

type Repository interface {
	AddURL(key, url string) error
	GetURL(key string) (string, error)
}
