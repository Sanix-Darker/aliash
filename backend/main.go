package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// routes
	router.GET("/i", InstallHandler)
	router.POST("/new", CreateAliasesHandler)
	router.GET("/:uid", GetHandler)
	router.GET("/search", SearchHandler)
	router.GET("/all", GetAllHandler)

	router.Run("localhost:" + os.Getenv("PORT"))
}
