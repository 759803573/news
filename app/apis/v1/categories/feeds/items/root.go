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

func item(c *gin.Context) *models.Item {
	item, ok := c.Get("item")

	if !ok {
		return &models.Item{}
	}

	return item.(*models.Item)
}

func getFeedItemsHandle(c *gin.Context) {
	paramsCategory, _ := c.Get("category")
	paramsFeed, _ := c.Get("feed")

	paramsItem := item(c)
	items := paramsFeed.(*models.Feed).GetItems(paramsItem, paramsCategory.(*models.Category).Feeds(nil))

	c.JSON(http.StatusOK, items)
}

func readFeedItemsHandle(c *gin.Context) {
	paramsCategory, _ := c.Get("category")
	paramsFeed, _ := c.Get("feed")

	paramsItem := item(c)
	items := paramsFeed.(*models.Feed).GetItems(paramsItem, paramsCategory.(*models.Category).Feeds(nil))
	if len(items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
	}
	items[0].Read()
}

func collectionFeedItemHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
