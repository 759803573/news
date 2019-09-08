package main

import (
	"news/app"
	"news/app/models"
	"news/config/initializers"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init()
	migrate()

	//workers.GetFeedsSyncWork().Run()

	app.New(":8080").Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
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
