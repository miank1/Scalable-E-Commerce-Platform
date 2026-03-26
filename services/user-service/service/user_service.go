package service

import (
	"database/sql"

	"user-service/repository"
)

func RegisterUser(db *sql.DB, name, email, password string) error {
	return repository.CreateUser(db, name, email, password)
}
