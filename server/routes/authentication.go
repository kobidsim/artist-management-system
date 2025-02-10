package routes

import (
	"artist-management-system/handler"
	"artist-management-system/service"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func SetupRegisterRoutes(app *echo.Echo, db *sql.DB) {
	service := service.NewAuthenticationService(db)
	handler := handler.NewAuthenticationHandler(service)

	app.POST("/login", handler.Login)
	app.POST("/register", handler.Register)
	app.GET("/logout", handler.Logout)
}
