package apis

import (
	v1 "news/app/apis/v1"

	"github.com/gin-gonic/gin"
)

//NewRoot API 挂载点
func NewRoot(g *gin.RouterGroup) {
	g1 := g.Group("/v1")
	v1.NewRoot(g1)
}
