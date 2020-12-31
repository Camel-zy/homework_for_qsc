package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

// the functions in this file are currently unused,
// and haven't been tested.

func GenerateJWT() (string, *time.Time) {
	secretKey := viper.GetString("jwt.secret_key")
	maxAge := 60 * 10     // 10 minuets
	expireTime := time.Now().Add(time.Duration(maxAge) * time.Second)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: "rop"})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	return tokenString, &expireTime
}

func ParseJWT(tokenString string) (*jwt.StandardClaims, error) {
	secretKey := viper.GetString("jwt.secret_key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpexted singing method: %v\n", token.Header["alg"])
		} else {
			return []byte(secretKey), nil
		}
	})
	if claims, ok := token.Claims.(jwt.StandardClaims); ok && token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}
