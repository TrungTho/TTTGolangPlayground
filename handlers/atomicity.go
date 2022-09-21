package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	freeValue = 0
	// redisLockValue = 0
)

func FreeAddValue(c *gin.Context) {
	freeValue++
}

func GetFreeValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"value": freeValue,
	})
}
