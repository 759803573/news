package categories

import (
	"fmt"
	"net/http"
	"news/app/apis/v1/categories/items"
	"news/app/models"

	"github.com/gin-gonic/gin"
)

//NewRoot API 挂载
func NewRoot(g *gin.RouterGroup) {
	g.GET("/", getCategoriesHandle)
	gCategory := g.Group(":category_id")
	gCategory.GET("/status", getCategoryStatusHandle)
	gItems := gCategory.Group("items")
	items.NewRoot(gItems)
}

func getCategoriesHandle(c *gin.Context) {
	category := models.Category{UserID: 1}
	if err := category.GetByUserID(); err != nil {
		fmt.Println("Not Found")
	}
	fmt.Println(category)
}

func getCategoryStatusHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func getCategoryItemsHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
