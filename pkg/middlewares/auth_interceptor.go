package middleware

import (
	"fmt"

	pathList "go-clean-arch/assets/path_list"
	"go-clean-arch/modules/entities/auth"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthInterceptor(next echo.HandlerFunc) echo.HandlerFunc {
	var jwtKey = []byte("one_wish")
	return func(c echo.Context) error {
		for _, path := range pathList.PublicPathConst() {
			if c.Request().RequestURI == path {
				return next(c)
			}
		}
		raw := c.Request().Header.Get("Authorization")
		tokenString := strings.TrimPrefix(raw, "Bearer ")
		fmt.Println(tokenString)

		// if(err := valida) {

		// }
		// return echo.ErrUnauthorized
		// Parse and verify the token
		claims := &auth.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
		}

		// Set claims in context for further use
		c.Set("user", claims)
		return next(c)

	}

}
