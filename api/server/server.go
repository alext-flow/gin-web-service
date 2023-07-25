package server

import (
	"api/server/endpoints"
	"api/server/db"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func (app *App) Initialize() error{
	app.router = gin.Default()

	//db connection
	err := db.SetupDatabaseConn()
	if err != nil {
		return err
	}
	
	return nil
}

func (app *App) ConfigureRoutes() {
	// app.router.GET("/type/:type", endpoints.GetAlbumType)
	app.router.GET("/albums", endpoints.GetAlbums)
	app.router.POST("/albums", endpoints.CreateAlbum)
	app.router.GET("/albums/:id", endpoints.GetAlbumByID)
	app.router.DELETE("/albums/:id", endpoints.DeleteAlbumById)
	// app.router.GET("/albums/type/:type", endpoints.GetAlbumsByType)
	app.router.PATCH("/albums/:id", endpoints.UpdateAlbum)
}

func (app *App) Serve() {
	app.router.Run()
}
