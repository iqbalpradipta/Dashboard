package middlewares

import (
	"fmt"
	"immersiveProject/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.SECRET_JWT),
	})
}

func CreateToken(UserId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["UserId"] = UserId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractToken(c echo.Context) (int, error) {
	headerData := c.Request().Header.Get("Authorization")
	dataAuth := strings.Split(headerData, " ")
	token := dataAuth[len(dataAuth)-1]
	datajwt, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_JWT), nil
	})

	if datajwt.Valid {
		claims := datajwt.Claims.(jwt.MapClaims)
		UserId := claims["UserId"].(float64)
		return int(UserId), nil
	}

	return -1, fmt.Errorf("token invalid")
}
func ExtractRoleToken(c echo.Context) (string, error) {
	headerData := c.Request().Header.Get("Authorization")
	dataAuth := strings.Split(headerData, " ")
	token := dataAuth[len(dataAuth)-1]
	datajwt, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_JWT), nil
	})

	if datajwt.Valid {
		claims := datajwt.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		return role, nil
	}

	return "", fmt.Errorf("token invalid")
}
