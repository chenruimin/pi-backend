package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Email     string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// Ping test
	r.POST("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.Email == "john@smith.com" && json.Password == "mypassword" {
			c.JSON(http.StatusOK, gin.H{"match": true})
		} else {
			c.JSON(http.StatusOK, gin.H{"match": false})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:9099
	r.Run(":9099")
}
