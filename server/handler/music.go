package handler

import (
	"artist-management-system/service"
	"artist-management-system/view"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type musicHandler struct {
	musicService service.MusicService
}

type MusicHandler interface {
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

func NewMusicHandler(musicService service.MusicService) MusicHandler {
	return musicHandler{
		musicService: musicService,
	}
}

func (handler musicHandler) List(ctx echo.Context) error {
	artistID, err := strconv.Atoi(ctx.QueryParam("artist_id"))
	if err != nil {
		fmt.Printf("ERROR:: invalid artist id: %s\n", err.Error())
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error": true,
			"data":  "Could not find artist",
		})
	}

	music, err := handler.musicService.All(artistID)
	if err != nil {
		fmt.Printf("ERROR:: error getting music: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": true,
			"data":  "Could not get music list",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": false,
		"data":  music,
	})
}

func (handler musicHandler) Create(ctx echo.Context) error {
	var params view.MusicView
	if err := ctx.Bind(&params); err != nil {
		fmt.Printf("ERROR:: error binding params: %s\n", err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": "Bad Request",
		})
	}

	if err := ctx.Validate(&params); err != nil {
		fmt.Printf("ERROR:: validation error: %s\n", err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": "Validation Error",
		})
	}

	if _, err := strconv.Atoi(params.ComposedByID); err != nil {
		fmt.Printf("ERROR:: invalid composer id: %s\n", err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": "Invalid composer id",
		})
	}

	if err := handler.musicService.Create(params); err != nil {
		fmt.Printf("ERROR:: error creating music: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Could not add music",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "Created music successfully",
	})
}

func (handler musicHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Printf("ERROR:: error getting id: %s\n", err.Error())
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   true,
			"message": "Music not found",
		})
	}

	var params view.UpdateMusicView
	if err := ctx.Bind(&params); err != nil {
		fmt.Printf("ERROR:: error binding params: %s\n", err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": "Bad request",
		})
	}

	if err := ctx.Validate(&params); err != nil {
		fmt.Printf("ERROR:: validation error: %s\n", err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": "Validation Error",
		})
	}

	if err := handler.musicService.Update(id, params); err != nil {
		fmt.Printf("ERROR:: error updating music: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Error editing music",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "Updated music successfully",
	})
}

func (handler musicHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Printf("ERROR:: error getting id: %s\n", err.Error())
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   true,
			"message": "Music not found",
		})
	}

	if err := handler.musicService.Delete(id); err != nil {
		fmt.Printf("ERROR:: error deleting artist: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Error deleting music",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "Deleted music successfully",
	})
}
