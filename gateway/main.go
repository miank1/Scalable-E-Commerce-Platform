package main

import (
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
