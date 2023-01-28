package service

import (
	"errors"
	"strings"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/internal/repository"
)

type PostService interface {
	LikePost(postId, userId int) error
	DislikePost(postId, userId int) error
	GetAllPosts() ([]model.Post, error)
	CreatePost(post *model.Post) (int64, error)
	CreateCategory(categories *model.Category) error
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

func (p *postService) CreatePost(post *model.Post) (int64, error) {
	if ok := validDataString(post.Title); !ok {
		return -1, errors.New("title is invalid")
	}
	if ok := validDataString(post.Content); !ok {
		return -1, errors.New("content is invalid")
	}
	if ok := validCategory(post.Category); !ok {
		return -1, errors.New("category is invalid")
	}

	id, err := p.repository.CreatePost(*post)
	if err != nil {
		return -500, errors.New("create post was failed")
	}
	return id, nil
}

func (p *postService) CreateCategory(categories *model.Category) error {
	return p.repository.CreateCategory(*categories)
}

func validDataString(s string) bool {
	str := strings.TrimSpace(s)
	for _, v := range str {
		if v < rune(32) {
			return false
		}
	}
	return true
}

func validCategory(s string) bool {
	str := strings.Split(s, " ")
	for _, v := range str {
		if v != "romance" && v != "adventure" && v != "comedy" && v != "drama" && v != "fantasy" {
			return false
		}
	}
	return true
}
