package handlers

import (
	"net/http"
	"playground/constants"
	"playground/redis"

	"github.com/gin-gonic/gin"
)

var (
	freeValue      = 0
	redisLockValue = 0
)

const Limit = 100

func FreeAddValue(c *gin.Context) {
	freeValue++
	if freeValue > Limit {
		freeValue--
	}

	c.Status(http.StatusOK)
}

func RedisLockAddValue(c *gin.Context) {
	if err := redis.LockWithKey(constants.KeyAtomicAddValue); err != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	redisLockValue++
	if redisLockValue > Limit {
		redisLockValue--
	}

	redis.ReleaseLockWithKey(constants.KeyAtomicAddValue)
	c.Status(http.StatusOK)
}

func ResetValue(c *gin.Context) {
	freeValue = 0
	redisLockValue = 0
	c.Status(http.StatusOK)
}

type GetValueResp struct {
	FreeValue      int `json:"free_value"`
	RedisLockValue int `json:"redis_lock_value"`
}

func GetFreeValue(c *gin.Context) {
	c.JSON(http.StatusOK, &GetValueResp{
		FreeValue:      freeValue,
		RedisLockValue: redisLockValue,
	})
}
