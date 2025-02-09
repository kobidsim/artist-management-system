package routes

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo, db *sql.DB) {
	SetupRegisterRoutes(app, db)
	SetupUserRoutes(app, db)
	SetupArtistRoutes(app, db)
}
