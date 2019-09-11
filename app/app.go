package app

import (
	"net/http"
	"news/app/apis"
	"news/app/models"
	"news/config"
	"time"

	"github.com/gin-gonic/gin"
)

//App App入口
type App struct {
	router     *gin.Engine
	httpServer *http.Server
}

//New Init App
func New(addr string) *App {
	app := &App{}

	app.inintServer(addr)

	return app
}

func (app *App) inintServer(addr string) {
	app.httpServer = &http.Server{
		Addr:           addr,
		Handler:        app.createRouter(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

//Run start api server
func (app *App) Run() {
	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()
}

//Term 停止服务
func (app *App) Term() {
}

func (app *App) createRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	apis.NewRoot(&r.RouterGroup)
	return r
}

//Migrate 执行 DB Migrate
func (app *App) Migrate() {
	config.DB.Conn.AutoMigrate(&models.User{})
	config.DB.Conn.AutoMigrate(&models.CategoryFeed{})
	config.DB.Conn.AutoMigrate(&models.Category{})
	config.DB.Conn.AutoMigrate(&models.Feed{})
	config.DB.Conn.AutoMigrate(&models.Item{})
	config.DB.Conn.AutoMigrate(&models.ItemStatus{})
	config.DB.Conn.AutoMigrate(&models.Collection{})
	user := &models.User{}
	user.ID = 1
	config.DB.Conn.FirstOrCreate(&user)
	category := &models.Category{Name: "Category1", UserID: user.ID}
	category.ID = 1
	config.DB.Conn.FirstOrCreate(category)
	category = &models.Category{Name: "Category2"}
	category.ID = 2
	config.DB.Conn.FirstOrCreate(category)

	categoryFeed := &models.CategoryFeed{CategoryID: 1, FeedID: 1}
	config.DB.Conn.FirstOrCreate(categoryFeed, categoryFeed)
	categoryFeed = &models.CategoryFeed{CategoryID: 2, FeedID: 2}
	config.DB.Conn.FirstOrCreate(categoryFeed, categoryFeed)
}
