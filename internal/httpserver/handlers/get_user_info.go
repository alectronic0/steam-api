package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"steam-api/internal/cache"
	"steam-api/internal/steamgamecomparator"
	"steam-api/internal/steamservice"
)

func GetUserInfo(service steamservice.IService, c *cache.RedisCache) func(*gin.Context) {
	return func(ctx *gin.Context) {
		userID1, _ := ctx.GetQuery("user_id_1")
		userID2, _ := ctx.GetQuery("user_id_2")

		// Create cache key
		cacheKey := fmt.Sprintf("comparison:%s:%s", userID1, userID2)

		// Try to get from cache
		var cachedResult interface{}
		if c != nil {
			if err := c.Get(context.Background(), cacheKey, &cachedResult); err == nil {
				ctx.JSON(200, cachedResult)
				return
			}
		}

		// Cache miss or cache disabled, compute result
		comparatorService := steamgamecomparator.New(service)
		result, err := comparatorService.CompareOwnedGames(userID1, userID2)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Store in cache for 1 hour
		if c != nil {
			if err := c.Set(context.Background(), cacheKey, result, 1*time.Hour); err != nil {
				fmt.Printf("Cache set error: %v\n", err)
				// Don't fail if cache write fails
			}
		}

		ctx.JSON(200, result)
	}
}
