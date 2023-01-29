package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/with-insomnia/Forum-Golang/internal/model"
)

// 400 - http.StatusBadRequest
// 500 - http.StatusInternalServerError
// 200 - http.StatusOk
func (p *postService) GetAllCommentsAndPostByPostID(id int64) (model.Post, int) {
	initialPost, err := p.repository.GetPostById(int64(id))
	if err != nil {
		log.Println(err)
		return model.Post{}, 400
	}
	comments, err := p.repository.GetAllCommentByPostID(int64(id))
	if err != nil {
		log.Println(err)
		return model.Post{}, 500
	}

	sortedComments := []model.Comment{}

	for i := len(comments) - 1; i >= 0; i-- {
		sortedComments = append(sortedComments, comments[i])
	}

	initialPost.Comment = sortedComments
	return initialPost, 200
}

func (p *postService) CreateComment(comment *model.Comment) (int, error) {
	if ok := validDataString(comment.Message); !ok {
		return 400, errors.New("comment message is invalid")
	}
	_, err := p.repository.GetPostById(comment.PostId)
	if err != nil {
		fmt.Println(err)
		return 400, errors.New("post not exist")
	}
	err = p.repository.CommentPost(*comment)
	if err != nil {
		return 500, errors.New("comment post was failed")
	}
	return 200, nil
}
