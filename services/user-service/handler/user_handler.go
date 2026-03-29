package handler

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"time"

	"user-service/model"
	"user-service/service"

	"github.com/gin-gonic/gin"
)

func Register(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		err := service.RegisterUser(db, user.Name, user.Email, user.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user created"})
	}
}

func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("LOGIN HIT")
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		token, err := service.LoginUser(db, user.Email, user.Password)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid credentials"})
			return
		}

		c.JSON(200, gin.H{
			"token": token,
		})
	}
}

func GetProductsFromProductService(c *gin.Context) {

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get("http://localhost:8081/products")

	if err != nil {
		fmt.Println("ERROR calling product service:", err)
		c.JSON(500, gin.H{"error": "failed to call product service"})
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR reading body:", err)
		c.JSON(500, gin.H{"error": "failed to read response"})
		return
	}

	fmt.Println("Response Body:", string(body))

	c.Data(200, "application/json", body)
}

func Profile(c *gin.Context) {

	fmt.Println("Profile HIT")

	email := c.GetHeader("X-User-Email")

	c.JSON(200, gin.H{
		"user": email,
	})
}
