package untils

import (
	"bookManagerSystem/api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

const token = "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZSwiaWQiOjEsImV4cCI6MTY1Mzc5MTg1MCwianRpIjoiand0X2lkIn0.t5XkGqk0ayGRyr43zngqxej-vyoXYtMEj-rHjYsl-97gWL2yUDvP-T8dHY5V2HuzOGFP4Vyscu1Tg8NY0jczwA"

var expClaims1 = middleware.JwtCustomClaims{
	Name:  "admin",
	Admin: true,
	Id:    1}

var expClaims2 = middleware.JwtCustomClaims{
	Name:  "admin1",
	Admin: false,
	Id:    2}

func TestEncodingUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ass := assert.New(t)
	claims, err := EncodingUser(c)
	if err != nil {
		t.Log("test is error")
	}
	ass.Equal(expClaims1.Id, claims.Id, "should equal")
	ass.Equal(expClaims1.Name, claims.Name, "should equal")
	ass.Equal(expClaims1.Admin, claims.Admin, "should equal")
	ass.NotEqual(expClaims2.Id, claims.Id)
	ass.NotEqual(expClaims2.Name, claims.Name)
	ass.NotEqual(expClaims2.Admin, claims.Admin)
}
