package main

import (
	"news/app/models"
	"news/config/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	migrate()
	r.Static("/assets", "./public/assets")
	r.Static("/views", "./app/views")

	r.Run()
}

func migrate() {
	initializers.DB.Init("", nil)
	initializers.DB.Conn.AutoMigrate(&models.User{})
	initializers.DB.Conn.AutoMigrate(&models.Category{})
	initializers.DB.Conn.AutoMigrate(&models.Feed{})
	initializers.DB.Conn.AutoMigrate(&models.Item{})
	initializers.DB.Conn.AutoMigrate(&models.ItemStatus{})
	initializers.DB.Conn.AutoMigrate(&models.Collection{})
}
