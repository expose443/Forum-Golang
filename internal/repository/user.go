package repository

import (
	"database/sql"

	"github.com/with-insomnia/Forum-Golang/internal/model"
)

type UserQuery interface {
	CreateUser(user *model.User) error
	GetUserIdByToken(token string) (int, error)
	GetUserByUserId(userID int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
}

type userQuery struct {
	db *sql.DB
}

func (u *userQuery) CreateUser(user *model.User) error {
	_, err := u.db.Exec("INSERT INTO users(email, username, password) values(?,?,?)", user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *userQuery) GetUserIdByToken(token string) (int, error) {
	row := u.db.QueryRow("SELECT user_id FROM sessions WHERE token=?", token)
	var userID int
	if err := row.Scan(&userID); err != nil {
		return -1, err
	}
	return userID, nil
}

func (u *userQuery) GetUserByUserId(userID int) (model.User, error) {
	row := u.db.QueryRow("SELECT user_id, email, password, username FROM users WHERE user_id = ?", userID)
	var user model.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Username); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userQuery) GetUserByEmail(email string) (model.User, error) {
	row := u.db.QueryRow("SELECT user_id,email,password,username FROM users WHERE email = ?", email)
	var user model.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Username); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userQuery) GetUserByUsername(username string) (model.User, error) {
	row := u.db.QueryRow("SELECT user_id,email,password,username FROM users WHERE username = ?", username)
	var user model.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Username); err != nil {
		return model.User{}, err
	}
	return user, nil
}
