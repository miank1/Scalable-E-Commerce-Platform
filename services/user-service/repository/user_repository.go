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
