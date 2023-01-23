package service

import "github.com/with-insomnia/Forum-Golang/internal/repository"

type AuthService interface{}

type authService struct {
	repository.SessionQuery
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{
		dao.NewSessionQuery(),
	}
}
