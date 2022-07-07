// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router := gin.Default()

	// Handle Index
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hi hello")
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
