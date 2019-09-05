package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./public/assets")
	r.Static("/views", "./app/views")
	r.Run()
}
