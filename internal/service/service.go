package service

import (
	"grpcprj/internal/model"
)

type Service struct {
	dbRepo model.Database
}

func NewService(dbRepo model.Database) model.Service {
	return &Service{
		dbRepo: dbRepo,
	}
}

func (s Service) GetUsers() (users string) {
	users = s.dbRepo.GetUsers()
	return users
}

func (s Service) SignUp(user model.User) error {
	err := s.dbRepo.SignUp(user)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Login(username string, password string) (res bool, err error) {
	login, err := s.dbRepo.Login(username, password)
	if err != nil {
		return false, err
	}
	return login, nil
}

func (s Service) Delete(username string) (res bool, err error) {
	del, err := s.dbRepo.Delete(username)
	if err != nil {
		return false, err
	}
	return del, err
}
