package middleware

import (
	"errors"

	pathList "go-clean-arch/assets/path_list"
	"go-clean-arch/modules/entities/auth"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthInterceptor(next echo.HandlerFunc) echo.HandlerFunc {
	var jwtKey = []byte("one_wish")
	const userClaims = "user"
	return func(c echo.Context) error {

		if pathList.IsPathPublic(c.Path()) {
			return next(c)
		}

		raw := c.Request().Header.Get("Authorization")
		tokenString := strings.TrimPrefix(raw, "Bearer ")

		// Parse and verify the token
		claims := &auth.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return jwtKey, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": "Token signature invalid",
				})
			}

			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": " Token Invalid",
			})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token Expire"})
		}

		// Set claims in context for further use
		c.Set(userClaims, claims)
		return next(c)

	}
}
