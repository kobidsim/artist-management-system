package handler

import (
	"artist-management-system/service"
	"artist-management-system/view"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
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
	var params view.UserView
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

func (handler userHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Printf("ERROR:: error getting id: %s\n", err.Error())
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   true,
			"message": "User not found",
		})
	}

	var params view.UserView
	if err := ctx.Bind(&params); err != nil {
		fmt.Printf("ERROR:: error binding params: %s\n", err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": "Bad request",
		})
	}

	if err := handler.userService.Update(id, params); err != nil {
		fmt.Printf("ERROR:: error updating user: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Error editing user",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "Updated user successfully",
	})
}
