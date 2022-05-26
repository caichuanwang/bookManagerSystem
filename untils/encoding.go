package untils

import (
	"bookManagerSystem/api/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func EncodingUser(c echo.Context) (*middleware.JwtCustomClaims, error) {
	token := c.Request().Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimSpace(strings.Trim(token, "Bearer"))
	}
	claims, err := middleware.ParseToken(token)
	if err != nil {
		return nil, c.JSON(http.StatusInternalServerError, err.Error())
	}
	return claims, nil
}
