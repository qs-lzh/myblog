package data

import (
	"database/sql"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

type UserModel struct {
	DB *sql.DB
}
