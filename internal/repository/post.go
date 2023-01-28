package repository

import (
	"database/sql"
	"fmt"

	"github.com/with-insomnia/Forum-Golang/internal/model"
)

type PostQuery interface {
	CreatePost(post model.Post) (int64, error)
	GetAllPost() ([]model.Post, error)
	GetPostById(postId int64) (model.Post, error)
	CreateCategory(category *model.Category) error
	GetCategory() ([]model.Category, error)
	GetDisLikeStatus(postId, userId int) int
	DeletePostDisLike(post_id, user_id int) error
	DisLikePost(post_id, user_id, status int) error
	GetLikedPostIdByUserId(userId int) ([]int64, error)
	GetLikeStatus(postId, userId int) int
	LikePost(post_id, user_id, status int) error
	UpdatePostLikeDislike(post_id, like, dislike int) error
	DeletePostLike(post_id, user_id int) error
	GetAllCommentByPostID(postId int64) ([]model.Comment, error)
	GetCommentByCommentID(commentId int64) (model.Comment, error)
	CommentPost(comment model.Comment) error

	GetCommentLikeStatus(comment_id, userId int) int
	LikeComment(comment_id, user_id, status int) error
	UpdateCommentLikeDislike(comment_id, like, dislike int) error
	DeleteCommentLike(comment_id, user_id int) error

	DisLikeComment(comment_id, user_id, status int) error
	DeleteCommentDisLike(comment_id, user_id int) error
	GetDisLikeCommentStatus(comment_id, userId int) int
}

type postQuery struct {
	db *sql.DB
}

func (p *postQuery) CreatePost(post model.Post) (int64, error) {
	res, err := p.db.Exec("INSERT INTO posts (title, message, user_id, username ,category, like, dislike, born) VALUES(?,?,?,?,?,?,?,?)", post.Title, post.Content, post.Author.ID, post.Author.Username, post.Category, 0, 0, post.CreateTime)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	post.ID = id
	return id, nil
}

func (p *postQuery) GetAllPost() ([]model.Post, error) {
	rows, err := p.db.Query("SELECT * FROM posts")
	if err != nil {
		return []model.Post{}, err
	}
	defer rows.Close()
	var all []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Author.ID, &post.Author.Username, &post.Title, &post.Content, &post.Like, &post.Dislike, &post.Category, &post.CreateTime); err != nil {
			return []model.Post{}, err
		}
		all = append(all, post)
	}
	return all, nil
}

func (p *postQuery) GetPostById(postId int64) (model.Post, error) {
	row := p.db.QueryRow("SELECT post_id, title, message, user_id, username, like, dislike, category, born FROM posts WHERE post_id = ? ", postId)
	var post model.Post
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Author.ID, &post.Author.Username, &post.Like, &post.Dislike, &post.Category, &post.CreateTime); err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (p *postQuery) CreateCategory(category *model.Category) error {
	query := `INSERT INTO categories(category, post_id) VALUES(?,?)`
	_, err := p.db.Exec(query, category.CategoryName, category.PostId)
	if err != nil {
		return err
	}
	return nil
}

func (p *postQuery) GetCategory() ([]model.Category, error) {
	query := `SELECT * FROM categories`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	var result []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.CategoryName, &category.PostId); err != nil {
			return nil, err
		}
		result = append(result, category)
	}
	return result, nil
}
