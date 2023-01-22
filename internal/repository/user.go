package repository

import "database/sql"

type UserQuery interface{}

type userQuery struct {
	db *sql.DB
}
