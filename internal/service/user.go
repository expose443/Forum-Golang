package service

import (
	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/internal/repository"
)

type UserService interface {
	GetUserByToken(token string) (*model.User, error)
}

type userService struct {
	repository repository.UserQuery
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{
		dao.NewUserQuery(),
	}
}

func (u *userService) GetUserByToken(token string) (*model.User, error) {
	userID, err := u.repository.GetUserIdByToken(token)
	if err != nil {
		return nil, err
	}

	user, err := u.repository.GetUserByUserId(*userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
