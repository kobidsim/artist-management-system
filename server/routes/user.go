package routes

import (
	"artist-management-system/handler"
	"artist-management-system/middleware"
	"artist-management-system/service"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(app *echo.Echo, db *sql.DB) {
	service := service.NewUserService(db)
	handler := handler.NewUserHandler(service)

	app.GET("/users", handler.List, middleware.AdminAuthMiddleware)
}
