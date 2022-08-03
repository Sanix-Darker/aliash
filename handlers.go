package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
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

	if len(as.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your content should not be empty",
		})
	}

	// We should check here if the user is logged, then allow content <= 1000
	if len(as.Content) > 300 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your content is > 300",
		})
	} else {
		as.Title = TruncateText(as.Title, 70)

		// We should search for the same title and the content in the database
		// and refuse to save if there is already something similar in the database
		filter := bson.M{"$and": []interface{}{bson.M{"title": as.Title}, bson.M{"content": as.Content}}}

		aliases, _ := filterAliasessBy(filter)

		if len(aliases) > 0 {
			c.JSON(http.StatusAlreadyReported, gin.H{
				"error": "An alias already exist for this title and this content",
			})
		} else {
			as.Hash512 = ShaIt(as.Content)
			as.Uid = TruncateText(slug.Make(as.Title), 7) + TruncateText(as.Hash512, 10)
			as.CreatedAt = time.Now()
			as.UpdatedAt = time.Now()

			if err := createAliases(&as); err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"uid":        as.Uid,
					"title":      as.Title,
					"hash":       as.Hash512,
					"created_at": as.CreatedAt,
				})
			}
		}

	}

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
			content := aliases[0].Content

			c.Data(http.StatusOK, "text/plain", []byte(content))
		}
	} else {
		log.Printf("[x] uid is empty !")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You need to provide the uid parameter !",
		})
	}
}
