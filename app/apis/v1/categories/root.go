package categories

import (
	"fmt"
	"net/http"
	"news/app/apis/v1/categories/items"
	"news/app/helpers"
	"news/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

const keyCategoryID = "category_id"

//NewRoot API 挂载
func NewRoot(g *gin.RouterGroup) {
	g.GET("/", getCategoriesHandle)
	gCategory := g.Group(fmt.Sprintf(":%s", keyCategoryID))
	gCategory.Use(helpers.MiddlewareHelpers.CheckFieldID(keyCategoryID))
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
	categoryID := c.GetString(keyCategoryID)
	category := &models.Category{}
	if categoryID != "*" {
		if categoryIDInt, err := strconv.Atoi(categoryID); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		} else {
			category.ID = uint(categoryIDInt)
		}
	}

	if categoryStatus := category.GetStatus(); categoryStatus != nil {
		c.JSON(http.StatusOK, categoryStatus)
	}
}

func getCategoryItemsHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
