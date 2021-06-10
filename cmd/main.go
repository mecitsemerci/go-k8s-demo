package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		name := ctx.Query("name")
		log.Printf("[NAME] = %s", name)
		ctx.JSON(http.StatusOK, gin.H{"Hello World!": name})
	})

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/readiness", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})

	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}
}
