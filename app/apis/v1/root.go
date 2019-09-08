package v1

import (
	"news/app/apis/v1/categories"

	"github.com/gin-gonic/gin"
)

//NewRoot API 挂载
func NewRoot(g *gin.RouterGroup) {
	gCategory := g.Group("/categories")
	categories.NewRoot(gCategory)
}
