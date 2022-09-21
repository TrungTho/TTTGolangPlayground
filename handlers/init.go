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

	router.Run(":" + os.Getenv("PORT"))

}
