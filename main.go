package main

import (
	"news/app"
	"news/config/initializers"
	"os"
	"os/signal"
	"path"
	"syscall"
)

func initializer() {
	initializers.DB.Init("", nil)
	return
}

func main() {
	initializer()

	// workers.GetFeedsSyncWork().Run()

	ap := app.New(":8080")
	ap.Migrate()
	ap.Static("/views", path.Join(".", "app", "views"))
	ap.Static("", path.Join(".", "app", "assets"))
	ap.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ap.Term()
}
