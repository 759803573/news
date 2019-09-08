package items

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//NewRoot API 挂载
func NewRoot(g *gin.RouterGroup) {
	g.GET("/", getCategoryItemsHandle)
	g.GET("/:item_id", getCategoryItemHandle)
	g.POST("/:item_id/read", readCategoryItemsHandle)
	g.POST("/:item_id/collection", collectionCategoryItemHandle)
}

func getCategoryItemsHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func getCategoryItemHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func readCategoryItemsHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func collectionCategoryItemHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
