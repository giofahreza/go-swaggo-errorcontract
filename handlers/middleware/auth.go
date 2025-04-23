package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Simulate authentication logic
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
			}

			// Remove bearer token prefix
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				authHeader = authHeader[7:]
			}

			// Simulate token validation
			if authHeader != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}
			return next(c)
		}
	}
}
