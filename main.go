package main

import (
	"news/app"
	"news/config/initializers"
	"os"
	"os/signal"
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
	ap.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
