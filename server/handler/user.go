package handler

import (
	"artist-management-system/service"
	"artist-management-system/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
}

func NewUserHandler(userService service.UserService) UserHandler {
	return userHandler{
		userService: userService,
	}
}

func (handler userHandler) List(ctx echo.Context) error {
	users, err := handler.userService.All()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": true,
			"data":  "Could not get users",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": false,
		"data":  users,
	})
}

func (handler userHandler) Create(ctx echo.Context) error {
	var params view.CreateUserView
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": "Bad Request",
		})
	}

	if err := handler.userService.Create(params); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Could not add user",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "Created user successfully",
	})
}
