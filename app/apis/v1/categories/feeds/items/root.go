package items

import (
	"fmt"
	"net/http"
	"news/app/helpers"
	"news/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

const keyItemID = models.KeyItemID

//NewRoot API 挂载
func NewRoot(g *gin.RouterGroup) {
	g.GET("/", getFeedItemsHandle)
	gItem := g.Group(fmt.Sprintf(":%s", keyItemID))
	gItem.Use(helpers.MiddlewareHelpers.CheckFieldID(keyItemID))
	gItem.Use(middlewareGenerateItem(keyItemID))
	gItem.GET("/", getFeedItemsHandle)
	gItem.POST("/read", readFeedItemsHandle)
	gItem.POST("/collection", collectionFeedItemHandle)
}

func middlewareGenerateItem(fieldName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString(fieldName)
		item := models.Item{}
		if id != "*" {
			if idInt, err := strconv.Atoi(id); err != nil {
				c.String(http.StatusBadRequest, err.Error())
			} else {
				item.ID = uint(idInt)
			}
		}
		c.Set("item", &item)
	}
}

func getFeedItemsHandle(c *gin.Context) {
	paramsCategory, _ := c.Get("category")
	paramsFeed, _ := c.Get("feed")
	//paramsFeed, _ := c.Get("feed")
	item, ok := c.Get("item")
	if !ok {
		item = &models.Item{}
	}
	items := paramsFeed.(*models.Feed).GetItems(item.(*models.Item), paramsCategory.(*models.Category).Feeds(nil))

	c.JSON(http.StatusOK, items)
}

func readFeedItemsHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func collectionFeedItemHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
