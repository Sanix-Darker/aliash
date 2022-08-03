package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// EmptyHandler ideal for HEAD calls
func EchoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "aliash is up and running !",
		"version": "1.0.1",
	})
}

func InstallHandler(c *gin.Context) {

	installScript, err := os.ReadFile("./install.sh")
	Must(err)

	c.Data(http.StatusOK, "text/plain", installScript)
}

func CreateAliasesHandler(c *gin.Context) {
	var as Aliases

	Must(c.BindJSON(&as))

	as.CreatedAt = time.Now()
	as.UpdatedAt = time.Now()

	if err := createAliases(&as); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":         as.Uid,
		"url":         as.Url,
		"description": as.Description,
		"created_at":  as.CreatedAt,
	})
}

func GetHandler(c *gin.Context) {
	uid := c.Param("uid")

	if len(uid) > 0 {
		filter := bson.D{{"uid", uid}}
		aliases, err := filterAliasessBy(filter)

		if err != nil || len(aliases) == 0 {

			log.Printf("%s", err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No aliases found for this uid !",
			})
		} else {
			url := aliases[0].Url

			if len(url) > 0 {
				rawScript := GetRequest(aliases[0].Url)

				c.Data(http.StatusOK, "text/plain", rawScript)
			} else {
				c.JSON(http.StatusNotAcceptable, gin.H{
					"error": "The url is not valid for this alias !",
				})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You need to provide the uid parameter !",
		})
	}
}
