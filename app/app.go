package app

import (
	"net/http"
	"news/app/apis"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	HTTPServer *http.Server
}

func New(addr string) *App {
	httpServer := &http.Server{
		Addr:           addr,
		Handler:        createRouter(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HTTPServer: httpServer,
	}
}

func (app *App) Run() {
	go func() {
		if err := app.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()
}

func createRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	apis.NewRoot(&r.RouterGroup)
	return r
}
