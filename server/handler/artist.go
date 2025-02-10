package handler

import (
	"artist-management-system/service"
	"artist-management-system/view"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type artistHandler struct {
	artistService service.ArtistService
}

type ArtistHandler interface {
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	CSVImport(ctx echo.Context) error
	CSVExport(ctx echo.Context) error
}

func NewArtistHandler(artistService service.ArtistService) ArtistHandler {
	return artistHandler{
		artistService: artistService,
	}
}

func (handler artistHandler) List(ctx echo.Context) error {
	artists, err := handler.artistService.All()
	if err != nil {
		fmt.Printf("ERROR:: error getting artists: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": true,
			"data":  "Could not get artists",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": false,
		"data":  artists,
	})
}

func (handler artistHandler) Create(ctx echo.Context) error {
	var params view.ArtistView
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

	dob, err := time.Parse("2006-01-02T15:04:05.000Z", params.Dob)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "dob is invalid")
	}
	if dob.After(time.Now().UTC()) {
		return ctx.JSON(http.StatusBadRequest, "dob can not be in the future")
	}

	if err := handler.artistService.Create(params); err != nil {
		fmt.Printf("ERROR:: error creating artist: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Could not add artist",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "Created artist successfully",
	})
}

func (handler artistHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Printf("ERROR:: error getting id: %s\n", err.Error())
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   true,
			"message": "Artist not found",
		})
	}

	var params view.ArtistView
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

	dob, err := time.Parse("2006-01-02T15:04:05.000Z", params.Dob)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "dob is invalid")
	}
	if dob.After(time.Now().UTC()) {
		return ctx.JSON(http.StatusBadRequest, "dob can not be in the future")
	}

	if err := handler.artistService.Update(id, params); err != nil {
		fmt.Printf("ERROR:: error updating artist: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Error editing artist",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "Updated artist successfully",
	})
}

func (handler artistHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Printf("ERROR:: error getting id: %s\n", err.Error())
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   true,
			"message": "Artist not found",
		})
	}

	if err := handler.artistService.Delete(id); err != nil {
		fmt.Printf("ERROR:: error deleting artist: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Error deleting artist",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "Deleted artist successfully",
	})
}

func (handler artistHandler) CSVImport(ctx echo.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil || form == nil {
		fmt.Printf("ERROR:: error getting form: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Bad request",
		})
	}

	fileHeader := form.File["csv_file"][0]
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Printf("ERROR:: error opening file: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Could not open file",
		})
	}
	defer file.Close()

	if err := handler.artistService.CSVImport(file); err != nil {
		fmt.Printf("ERROR:: error reading file: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Error reading file. File may not be a csv",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "File accepted",
	})
}

func (handler artistHandler) CSVExport(ctx echo.Context) error {
	csvBuffer, err := handler.artistService.CSVExport()
	if err != nil {
		fmt.Printf("ERROR:: error exporting: %s\n", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Error exporting csv",
		})
	}

	ctx.Response().Header().Set("Content-Disposition", "attachment; filename=artists.csv")
	return ctx.Blob(http.StatusOK, "text/csv", csvBuffer.Bytes())
}
