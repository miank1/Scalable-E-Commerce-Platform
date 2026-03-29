package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// USER SERVICE
	r.Any("/api/users/*path", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8080", "/api/users")
	})

	// PRODUCT SERVICE
	r.Any("/api/products/*path", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8081", "/api/products")
	})

	r.Run(":8000") // Gateway port
}

func forwardRequest(c *gin.Context, target string, prefix string) {
	path := strings.TrimPrefix(c.Request.URL.Path, prefix)

	url := target + path

	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "request creation failed"})
		return
	}

	req.Header = c.Request.Header

	email, exists := c.Get("user_email")
	if exists {
		req.Header.Set("X-User-Email", email.(string))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "service unavailable"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Middleware HIT ---------------->")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// ✅ REAL JWT PARSING
		email, err := extractEmailFromJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// ✅ SET EMAIL IN CONTEXT
		c.Set("user_email", email)

		c.Next()
	}
}
