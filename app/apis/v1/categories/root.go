package categories

import (
	"fmt"
	"strconv"

	"net/http"
	"news/app/helpers"
	"news/app/models"
	feeds "news/app/apis/v1/categories/feeds"

	"github.com/gin-gonic/gin"
)

const keyCategoryID = models.KeyCategoryID

//NewRoot /Categories 挂载点
func NewRoot(g *gin.RouterGroup) {
	g.GET("/", getCategoriesHandle)
	gCategory := g.Group(fmt.Sprintf(":%s", keyCategoryID))
	gCategory.Use(helpers.MiddlewareHelpers.CheckFieldID(keyCategoryID))
	gCategory.Use(middlewareGenerateCategory(keyCategoryID))
	gCategory.GET("/status", getCategoryStatusHandle)
	gFeeds := gCategory.Group("feeds")
	feeds.NewRoot(gFeeds)
}

func middlewareGenerateCategory(fieldName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString(fieldName)
		category := models.Category{UserID: 1}
		if id != "*" {
			if categoryIDInt, err := strconv.Atoi(id); err != nil {
				c.String(http.StatusBadRequest, err.Error())
			} else {
				category.ID = uint(categoryIDInt)
			}
		}
		c.Set("category", &category)
	}
}

func getCategoriesHandle(c *gin.Context) {
	category := models.Category{UserID: 1}
	if err := category.GetByUserID(); err != nil {
		fmt.Println("Not Found")
	}
	fmt.Println(category)
}

func getCategoryStatusHandle(c *gin.Context) {
	category, _ := c.Get("category")
	fmt.Println(category)

	if categoryStatus := category.(*models.Category).GetStatus(); categoryStatus != nil {
		c.JSON(http.StatusOK, categoryStatus)
	}
}
