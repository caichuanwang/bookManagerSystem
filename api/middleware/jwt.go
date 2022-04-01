package middleware

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

type JwtCustomClaims struct {
	Name  string `json:"name" form:"name"`
	Admin bool   `json:"admin" form:"admin"`
	jwt.StandardClaims
}

func CreateToken(name string, admin bool) (map[string]interface{}, error) {
	claims := JwtCustomClaims{
		name,
		admin,
		jwt.StandardClaims{
			Id:        "jwt_id",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and return.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return map[string]interface{}{}, err
	}
	return echo.Map{"token": t}, nil
}
