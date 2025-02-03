package main

import (
	"artist-management-system/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	db, err := database.NewDatabase()
	if err != nil {
		app.Logger.Fatal(err.Error())
	}
	defer db.Close()

	app.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello World")
	})

	app.Logger.Fatal(app.Start(":8080"))
}
