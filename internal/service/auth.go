package service

import (
	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/internal/repository"
)

type AuthService interface {
	Login(user *model.User, session *model.Session) (*model.Session, error)
	Register(user *model.User) error
	Logout(token string) error
}

type authService struct {
	repository.SessionQuery
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{
		dao.NewSessionQuery(),
	}
}

func (a *authService) Login(user *model.User, session *model.Session) (*model.Session, error) {
	return nil, nil
}

func (a *authService) Register(user *model.User) error {
	return nil
}

func (a *authService) Logout(token string) error {
	return nil
}
