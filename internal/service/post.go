package service

import "github.com/with-insomnia/Forum-Golang/internal/repository"

type PostService interface{}

type postService struct {
	repository repository.PostQuery
}

func NewPostService(dao repository.DAO) PostService {
	return &postService{
		dao.NewPostQuery(),
	}
}
