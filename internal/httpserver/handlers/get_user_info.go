package handlers

import (
	"github.com/gin-gonic/gin"
	"steam-api/internal/steamgamecomparator"
	"steam-api/internal/steamservice"
)

func GetUserInfo(service steamservice.IService) func(*gin.Context) {
	return func(c *gin.Context) {
		userID1, _ := c.GetQuery("user_id_1")
		userID2, _ := c.GetQuery("user_id_2")

		comparatorService := steamgamecomparator.New(service)
		result, err := comparatorService.CompareOwnedGames(userID1, userID2)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, result)
		return
	}
}
