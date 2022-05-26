package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

type JwtCustomClaims struct {
	Name  string `json:"name" form:"name"`
	Admin bool   `json:"admin" form:"admin"`
	Id    int64  `json:"id" form:"id"`
	jwt.StandardClaims
}

func CreateToken(name string, admin bool, id int64) (map[string]interface{}, error) {
	claims := JwtCustomClaims{
		name,
		admin,
		id,
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

func ParseToken(token string) (*JwtCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	claim := tokenClaims.Claims.(*JwtCustomClaims)
	return claim, nil
}
