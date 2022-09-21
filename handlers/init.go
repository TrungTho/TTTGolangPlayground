package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup() {
	router := gin.Default()

	//health-check
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	atomicAPI := router.Group("/atomic")
	{
		atomicAPI.GET("/free-add", FreeAddValue)
		atomicAPI.GET("/redis-lock-add", RedisLockAddValue)
		atomicAPI.GET("/reset-value", ResetValue)
		atomicAPI.GET("/value", GetFreeValue)
	}

	router.Run(":" + os.Getenv("PORT"))
}
