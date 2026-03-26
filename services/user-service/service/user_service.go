package service

import (
	"database/sql"
	"user-service/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, name, email, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return repository.CreateUser(db, name, email, string(hashedPassword))
}
