package handler

import (
	"artist-management-system/service"
	"artist-management-system/view"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authenticationHandler struct {
	authenticationService service.AuthenticationService
}

type AuthenticationHandler interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
	Logout(ctx echo.Context) error
}

func NewAuthenticationHandler(authenticationService service.AuthenticationService) AuthenticationHandler {
	return authenticationHandler{
		authenticationService: authenticationService,
	}
}

func (handler authenticationHandler) Login(ctx echo.Context) error {
	var params view.LoginView
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	response, err := handler.authenticationService.Login(params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Username or password incorrect")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":  response,
		"error": false,
	})
}

func (handler authenticationHandler) Register(ctx echo.Context) error {
	var params view.RegisterView
	if err := ctx.Bind(&params); err != nil {
		fmt.Printf("%s", err.Error())
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	if err := ctx.Validate(params); err != nil {
		fmt.Printf("%s", err.Error())
		return ctx.JSON(http.StatusBadRequest, "Validation error")
	}

	if err := handler.authenticationService.Register(params); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}

func (handler authenticationHandler) Logout(ctx echo.Context) error {
	return nil
}
