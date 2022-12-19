package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Generate(email, role string) (string, error) {

	secret := os.Getenv("SECRET")

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(12 * time.Hour).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ExtractEmailRole(c *gin.Context) (string, string, error) {

	secret := os.Getenv("SECRET")

	tokenString := extract(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		email := claims["email"].(string)
		role := claims["role"].(string)
		return email, role, nil
	}

	return "", "", nil
}

func extract(c *gin.Context) string {

	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}
