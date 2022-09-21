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

const Limit = 10

func FreeAddValue(c *gin.Context) {
	if freeValue+1 <= Limit {
		freeValue++
		c.String(http.StatusOK, "OK")
		return
	}

	c.AbortWithStatus(http.StatusBadRequest)
}

func RedisLockAddValue(c *gin.Context) {
	if err := redis.LockWithKey(constants.KeyAtomicAddValue); err != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	if freeValue+1 <= Limit {
		freeValue++
		c.String(http.StatusOK, "OK")
		redis.ReleaseLockWithKey(constants.KeyAtomicAddValue)
		return
	}

	redis.ReleaseLockWithKey(constants.KeyAtomicAddValue)
	c.AbortWithStatus(http.StatusBadRequest)
}

func GetFreeValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"free_value":         freeValue,
		"redis_locked_value": redisLockValue,
	})
}
