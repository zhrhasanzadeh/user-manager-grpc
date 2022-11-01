package model

type Service interface {
	GetUsers() (users string)
	SignUp(user User) error
	Login(username string, password string) (res bool, err error)
	Delete(username string) (res bool, err error)
}
