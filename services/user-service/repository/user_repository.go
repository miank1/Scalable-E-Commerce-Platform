package repository

import (
	"database/sql"
	"errors"
	"strings"
)

func CreateUser(db *sql.DB, name, email, password string) error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, name, email, password)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("email already exists")
		}
		return err
	}
	return err
}

func GetUserByEmail(db *sql.DB, email string) (string, error) {
	var hashedPassword string

	query := `SELECT password FROM users WHERE email=$1`

	err := db.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}

	return hashedPassword, nil
}
