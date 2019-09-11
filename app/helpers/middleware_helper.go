package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//MiddlewareHelpers middleware manager
var MiddlewareHelpers = &middlewareHelper{}

type middlewareHelper struct{}

func (middleware *middlewareHelper) CheckFieldID(fieldName string) gin.HandlerFunc {
	if fieldName == "" {
		panic(fmt.Errorf("fiedlName Not Allow Empty"))
	}

	return func(c *gin.Context) {
		if fieldID, ok := c.Params.Get(fieldName); ok {
			c.Set(fieldName, fieldID)
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("%s Not Found", fieldName)})
		}
	}
}
