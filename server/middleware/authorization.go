package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("ERROR:: invalid auth header: ", authHeader)
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   true,
				"message": "Unauthorized",
			})
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			fmt.Printf("ERROR:: error parsing token: %s\n", err.Error())
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   true,
				"message": "Unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("ERROR:: token does not have claims")
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   true,
				"message": "Invalid auth token",
			})
		}

		role, ok := claims["role"].(string)
		if !ok {
			fmt.Println("ERROR:: token does not have role claim")
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   true,
				"message": "Invalid auth token",
			})
		}

		if role != "super_admin" {
			fmt.Println("ERROR:: user does not have admin priviledges")
			return ctx.JSON(http.StatusForbidden, map[string]interface{}{
				"error":   true,
				"message": "Must be an admin user",
			})
		}

		ctx.Set("requestedByUserID", claims["id"])

		return next(ctx)
	}
}

func AdminManagerAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("ERROR:: invalid auth header: ", authHeader)
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   true,
				"message": "Unauthorized",
			})
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			fmt.Printf("ERROR:: error parsing token: %s\n", err.Error())
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   true,
				"message": "Unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("ERROR:: token does not have claims")
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   true,
				"message": "Invalid auth token",
			})
		}

		role, ok := claims["role"].(string)
		if !ok {
			fmt.Println("ERROR:: token does not have role claim")
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   true,
				"message": "Invalid auth token",
			})
		}

		if role != "super_admin" && role != "artist_manager" {
			fmt.Println("ERROR:: user does not have admin or artist_manager priviledges")
			return ctx.JSON(http.StatusForbidden, map[string]interface{}{
				"error":   true,
				"message": "Must be an admin or artist manager user",
			})
		}

		ctx.Set("requestedByUserID", claims["id"])

		return next(ctx)
	}
}
