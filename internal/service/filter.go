package service

import (
	"strings"

	"github.com/with-insomnia/Forum-Golang/internal/model"
)

// 400 - http status Bad request
// 500 - http status Internal server error
// 200 - http status Ok

func (p *postService) GetFilterPosts(genre string, user model.User) (model.Data, int) {
	if ok := validCategoryFilter(genre); !ok {
		return model.Data{}, 400
	}
	var posts []model.Post
	switch genre {
	case "liked-post":
		postId, _ := p.repository.GetLikedPostIdByUserId(int(user.ID))
		for _, v := range postId {
			post, err := p.repository.GetPostById(v)
			if err != nil {
				return model.Data{}, 400
			}
			posts = append(posts, post)
		}
		data := model.Data{
			Posts: posts,
			User:  user,
			Genre: "liked-post",
		}
		return data, 200
	case "created-post":
		allPost, _ := p.GetAllPosts()
		for _, v := range allPost {
			if v.Author.ID == user.ID {
				posts = append(posts, v)
			}
		}
		data := model.Data{
			Posts: posts,
			User:  user,
			Genre: "created-post",
		}
		return data, 200
	default:
		categories, err := p.repository.GetCategory()
		if err != nil {
			return model.Data{}, 500
		}
		var postId []int64
		for _, v := range categories {
			category := strings.Fields(v.CategoryName)
			for _, k := range category {
				if k == genre {
					postId = append(postId, v.PostId)
					break
				}
			}
		}
		for _, v := range postId {
			post, err := p.repository.GetPostById(v)
			if err != nil {
				return model.Data{}, 500
			}
			posts = append(posts, post)
		}

		data := model.Data{
			Posts: posts,
			User:  user,
			Genre: genre,
		}
		return data, 200
	}
}

func (p *postService) GetWelcomeFilterPosts(genre string) (model.Data, int) {
	var posts []model.Post
	categories, err := p.repository.GetCategory()
	if err != nil {
		return model.Data{}, 500
	}
	var postId []int64
	for _, v := range categories {
		category := strings.Fields(v.CategoryName)
		for _, k := range category {
			if k == genre {
				postId = append(postId, v.PostId)
				break
			}
		}
	}
	for _, v := range postId {
		post, err := p.repository.GetPostById(v)
		if err != nil {
			return model.Data{}, 500
		}
		posts = append(posts, post)
	}

	data := model.Data{
		Posts: posts,
		Genre: genre,
	}
	return data, 200
}
