package service

import (
	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/internal/repository"
)

type PostService interface {
	LikePost(postId, userId int) error
	DislikePost(postId, userId int) error
	GetAllPosts() ([]model.Post, error)
}

type postService struct {
	repository repository.PostQuery
}

func NewPostService(dao repository.DAO) PostService {
	return &postService{
		dao.NewPostQuery(),
	}
}

func (p *postService) GetAllPosts() ([]model.Post, error) {
	posts, err := p.repository.GetAllPost()
	if err != nil {
		return nil, err
	}
	result := []model.Post{}
	for i := len(posts) - 1; i >= 0; i-- {
		result = append(result, posts[i])
	}
	return result, nil
}
