package repository

import "database/sql"

type PostQuery interface{}

type postQuery struct {
	db *sql.DB
}
