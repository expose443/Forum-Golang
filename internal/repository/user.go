package repository

import (
	"database/sql"

	"github.com/with-insomnia/Forum-Golang/internal/model"
)

type UserQuery interface {
	GetUserIdByToken(token string) (int, error)
	GetUserByUserId(userID int) (model.User, error)
}

type userQuery struct {
	db *sql.DB
}

func (u *userQuery) GetUserIdByToken(token string) (int, error) {
	stmt, err := u.db.Prepare("SELECT user_id FROM sessions WHERE token=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(token)
	var userID int
	if err := row.Scan(&userID); err != nil {
		return 0, nil
	}
	return userID, nil
}

func (u *userQuery) GetUserByUserId(userID int) (model.User, error) {
	stmt, err := u.db.Prepare("SELECT * FROM users WHERE user_id = ?")
	if err != nil {
		return model.User{}, err
	}
	row := stmt.QueryRow(userID)
	var user model.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Username); err != nil {
		return model.User{}, err
	}
	return user, nil
}
