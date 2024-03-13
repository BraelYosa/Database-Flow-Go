package connection

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "No token provided",
			})
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "Invalid token format",
			})
		}

		tokenString = parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JwtSKey, nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
		}

		c.Set("admin_mail", claims["admin_mail"])
		c.Set("admin_name", claims["admin_name"])

		return next(c)
	}
}
