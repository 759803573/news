package Feeds

import (
	"fmt"
	"net/http"
	"news/app/apis/v1/categories/feeds/items"
	"news/app/helpers"
	"news/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

const keyFeedID = models.KeyFeedID

//NewRoot API 挂载
func NewRoot(g *gin.RouterGroup) {
	g.GET("/", getCategoryFeedsHandle)
	gFeed := g.Group(fmt.Sprintf(":%s", keyFeedID))
	gFeed.Use(helpers.MiddlewareHelpers.CheckFieldID(keyFeedID))
	gItem := gFeed.Group("items")
	items.NewRoot(gItem)
}

func middlewareGenerateFeed(fieldName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString(fieldName)
		feed := models.Feed{}
		if id != "*" {
			if feedIDInt, err := strconv.Atoi(id); err != nil {
				c.String(http.StatusBadRequest, err.Error())
			} else {
				feed.ID = uint(feedIDInt)
			}
		}
		c.Set("feed", &feed)
	}
}

func getCategoryFeedsHandle(c *gin.Context) {
	category, _ := c.Get("category")
	fmt.Println(category.(models.Category))
	c.String(http.StatusOK, "pong")
}
