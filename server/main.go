package main

import (
	"artist-management-system/database"
	"artist-management-system/routes"
	custom_validator "artist-management-system/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	db, err := database.NewDatabase()
	if err != nil {
		app.Logger.Fatal(err.Error())
	}
	defer db.Close()
	app.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	app.Validator = custom_validator.NewCustomValidator()

	routes.SetupRoutes(app, db)

	app.Logger.Fatal(app.Start(":8080"))
}
