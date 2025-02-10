package routes

import (
	"artist-management-system/handler"
	"artist-management-system/service"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func SetupMusicRoutes(app *echo.Echo, db *sql.DB) {
	service := service.NewMusicService(db)
	handler := handler.NewMusicHandler(service)

	app.GET("/music", handler.List)
	app.POST("/music", handler.Create)
	app.POST("/music/:id", handler.Update)
	app.DELETE("/music/:id", handler.Delete)
}
