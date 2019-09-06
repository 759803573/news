package main

import (
	"fmt"
	"news/app/models"
	"news/config"
	"news/config/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initializers.DB.Init("", nil)
	x := models.User{}
	initializers.DB.Conn.AutoMigrate(&x)
	fmt.Println(config.DB, x)
	r.Static("/assets", "./public/assets")
	r.Static("/views", "./app/views")
	r.Run()
}
