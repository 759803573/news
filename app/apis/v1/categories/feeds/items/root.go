package items

import (
	"fmt"
	"net/http"
	"news/app/helpers"
	"news/app/models"

	"github.com/gin-gonic/gin"
)

const keyItemID = models.KeyItemID

//NewRoot API 挂载
func NewRoot(g *gin.RouterGroup) {
	g.GET("/", nil)
	gFeed := g.Group(fmt.Sprintf(":%s", keyItemID))
	gFeed.Use(helpers.MiddlewareHelpers.CheckFieldID(keyItemID))
	gFeed.Group("items")
	gFeed.GET("/", getFeedItemsHandle)
	gFeed.POST("/read", readFeedItemsHandle)
	gFeed.POST("/collection", collectionFeedItemHandle)
}

func getFeedItemsHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func readFeedItemsHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func collectionFeedItemHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
