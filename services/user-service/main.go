package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Get DB connection string from env
	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB ping error:", err)
	}

	fmt.Println("✅ DB connected")

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/register", func(c *gin.Context) {
		var user User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
	`

		_, err := db.Exec(query, user.Name, user.Email, user.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to insert user"})
			return
		}

		c.JSON(200, gin.H{"message": "user created"})
	})

	r.Run(":8080")
}
