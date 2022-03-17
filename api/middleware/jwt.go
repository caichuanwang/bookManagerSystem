package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

type JwtCustomClaims struct {
	Name  string `json:"name" form:"name"`
	Admin bool   `json:"admin" form:"admin"`
	jwt.StandardClaims
}

func CreateToken(name string, admin bool) (map[string]interface{}, error) {
	claims := &JwtCustomClaims{
		name,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and return.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return map[string]interface{}{}, err
	}
	return echo.Map{"token": t}, nil
}
