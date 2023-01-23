package repository

import (
	"database/sql"
	"errors"

	"github.com/with-insomnia/Forum-Golang/internal/model"
)

type SessionQuery interface {
	GetSessionByToken(token string) (*model.Session, error)
	GetSessionByUserID(userId int) (*model.Session, error)
	CreateSession(session model.Session) error
	DeleteSession(token string) error
}

type sessionQuery struct {
	db *sql.DB
}

func (s *sessionQuery) CreateSession(session model.Session) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO sessions(email, username, token, expiry) VALUES(?,?,?,?)", session.Email, session.Username, session.Token, session.Expiry)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *sessionQuery) GetSessionByToken(token string) (*model.Session, error) {
	stmt, err := s.db.Prepare("SELECT session_id, email, username, token, expiry FROM sessions WHERE token = ?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(token)
	var session model.Session
	if err := row.Scan(&session.ID, &session.Email, &session.Username, &session.Token, &session.Expiry); err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *sessionQuery) GetSessionByUserID(userId int) (*model.Session, error) {
	stmt, err := s.db.Prepare("SELECT session_id, user_id,  email, username, token, expiry FROM sessions WHERE user_id = ?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(userId)
	var session model.Session
	if err := row.Scan(&session.ID, &session.UserId, &session.Email, &session.Username, &session.Token, &session.Expiry); err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *sessionQuery) DeleteSession(token string) error {
	stmt, err := s.db.Prepare("DELETE FROM sessions WHERE token = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(token)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("delete session was failed")
	}
	return nil
}
