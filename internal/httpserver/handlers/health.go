package handlers

import "github.com/gin-gonic/gin"

func Health() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	}
}
