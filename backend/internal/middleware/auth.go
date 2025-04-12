package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func SupabaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			fmt.Println("❌ Missing or invalid token format")
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		secret := os.Getenv("SUPABASE_JWT_SECRET")
		if secret == "" {
			fmt.Println("❌ SUPABASE_JWT_SECRET is missing!")
			return echo.NewHTTPError(http.StatusInternalServerError, "JWT secret not set")
		}
		fmt.Println("🔑 Using JWT secret:", secret)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			fmt.Println("❌ JWT parsing error:", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		if !token.Valid {
			fmt.Println("❌ Token is invalid")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)
		fmt.Println("✅ Token validated for user:", userID)

		c.Set("userID", userID)
		return next(c)
	}
}
