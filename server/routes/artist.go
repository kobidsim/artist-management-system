package routes

import (
	"artist-management-system/handler"
	"artist-management-system/middleware"
	"artist-management-system/service"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func SetupArtistRoutes(app *echo.Echo, db *sql.DB) {
	service := service.NewArtistService(db)
	handler := handler.NewArtistHandler(service)

	app.GET("/artists", handler.List)
	app.POST("/artist", handler.Create, middleware.AdminManagerAuthMiddleware)
	app.POST("/artist/:id", handler.Update, middleware.AdminManagerAuthMiddleware)
	app.DELETE("/artist/:id", handler.Delete, middleware.AdminManagerAuthMiddleware)
}
