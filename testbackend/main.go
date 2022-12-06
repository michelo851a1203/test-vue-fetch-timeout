package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Add("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Type, Content-Length, Cache-Control, X-CSRF-Token, X-Requested-With, Origin")
		ctx.Writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery(), CorsMiddleware())
	router.GET("/", func(ctx *gin.Context) {
		time.Sleep(1 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "testing",
		})
	})

	router.Run()
}
