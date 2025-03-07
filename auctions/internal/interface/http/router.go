package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lioarce01/auction-microservices/internal/infrastructure/database"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	_, err := database.NewDBConnection()
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return r
}
