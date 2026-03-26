package repository

import "database/sql"

func CreateUser(db *sql.DB, name, email, password string) error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, name, email, password)
	return err
}
