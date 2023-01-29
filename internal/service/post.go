package service

import (
	"errors"
	"strings"

	"github.com/with-insomnia/Forum-Golang/internal/model"
	"github.com/with-insomnia/Forum-Golang/internal/repository"
)

type PostService interface {
	LikePost(postId, userId int) int
	DislikePost(postId, userId int) int
	LikeComment(commentId, userId int) int
	DisLikeComment(commentId, userId int) int
	GetAllPosts() ([]model.Post, error)
	CreatePost(post *model.Post) (int, error)
	GetAllCommentsAndPostByPostID(id int64) (model.Post, int)
	CreateComment(comment *model.Comment) (int, error)
	GetFilterPosts(genre string, user model.User) (model.Data, int)
	GetWelcomeFilterPosts(genre string) (model.Data, int)
}

type postService struct {
	repository repository.PostQuery
}

func NewPostService(dao repository.DAO) PostService {
	return &postService{
		dao.NewPostQuery(),
	}
}

// 400 - http status Bad request
// 500 - http status Internal server error
// 200 - http status Ok

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

func (p *postService) CreatePost(post *model.Post) (int, error) {
	if ok := validDataString(post.Title); !ok {
		return 400, errors.New("title is invalid")
	}
	if ok := validDataString(post.Content); !ok {
		return 400, errors.New("content is invalid")
	}
	if ok := validCategory(post.Category); !ok {
		return 400, errors.New("category is invalid")
	}

	id, err := p.repository.CreatePost(*post)
	if err != nil {
		return 500, errors.New("create post was failed")
	}

	categories := model.Category{
		CategoryName: post.Category,
		PostId:       id,
	}

	err = p.repository.CreateCategory(&categories)
	if err != nil {
		return 500, errors.New("create category was failed")
	}
	return 200, nil
}

func validDataString(s string) bool {
	str := strings.TrimSpace(s)
	if len(str) == 0 {
		return false
	}
	for _, v := range str {
		if v < rune(32) {
			return false
		}
	}
	return true
}

func validCategory(s string) bool {
	category := make(map[string]struct{})
	valid := []string{
		"romance",
		"adventure",
		"comedy",
		"drama",
		"fantasy",
	}
	for _, v := range valid {
		category[v] = struct{}{}
	}

	str := strings.Split(s, " ")
	for _, v := range str {
		if _, ok := category[v]; !ok {
			return false
		}
	}
	return true
}

func validCategoryFilter(s string) bool {
	category := make(map[string]struct{})
	valid := []string{
		"romance",
		"adventure",
		"comedy",
		"drama",
		"fantasy",
		"liked-post",
		"created-post",
	}
	for _, v := range valid {
		category[v] = struct{}{}
	}
	if _, ok := category[s]; !ok {
		return false
	}

	return true
}
