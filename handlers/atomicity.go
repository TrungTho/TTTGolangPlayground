package handlers

import (
	"math/rand"
	"net/http"
	"playground/constants"
	"playground/redis"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	freeValue      = 0
	redisLockValue = 0
)

const Limit = 100

func FreeAddValue(c *gin.Context) {
	if freeValue+1 <= Limit {
		time.Sleep(time.Millisecond * time.Duration(1000*rand.Float64()))
		freeValue++
		c.String(http.StatusOK, "OK")
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func RedisLockAddValue(c *gin.Context) {
	if err := redis.LockWithKey(constants.KeyAtomicAddValue); err != nil {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	if freeValue+1 <= Limit {
		time.Sleep(time.Millisecond * time.Duration(1000*rand.Float64()))
		freeValue++
		c.String(http.StatusOK, "OK")
		redis.ReleaseLockWithKey(constants.KeyAtomicAddValue)
		return
	}

	redis.ReleaseLockWithKey(constants.KeyAtomicAddValue)
	c.AbortWithStatus(http.StatusOK)
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
