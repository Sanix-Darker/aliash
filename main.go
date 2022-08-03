package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// routes
	router.GET("/", InstallHandler)
	router.POST("/new", CreateAliasesHandler)
	router.GET("/:uid", GetHandler)

	router.Run("localhost:" + os.Getenv("PORT"))
}
