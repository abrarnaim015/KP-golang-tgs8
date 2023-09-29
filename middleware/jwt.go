package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


type jwtCustomClaims struct {
	UserId  int `json:"userId"`
	Name string   `json:"bool"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId int, name string) string {
	var claims = jwtCustomClaims {
		userId,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resultJWT, _ := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	return resultJWT
}

func VerifyTokenJWT(tokenString string) (int, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
			return 0, "", err
	}
	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
			// Token valid, Anda dapat mengakses klaim yang ada
			return claims.UserId, claims.Name, nil
	} else {
			// Token tidak valid
			return 0, "", errors.New("Invalid token")
	}
}