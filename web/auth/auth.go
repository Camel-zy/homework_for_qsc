package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

// the functions in this file are currently unused.

func GenerateJWT(uid string) (string, *time.Time) {
	mySigningKey := []byte(viper.GetString("jwt.secret_key"))

	maxAge := 60 * 10     // 10 minuets
	expireTime := time.Now().Add(time.Duration(maxAge) * time.Second)

	claims := &jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "rop",
		Subject:   uid,      // FIXME: Encrypt this!!!
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return tokenString, &expireTime
}

func ParseJWT(tokenString string) *jwt.Token {
	mySigningKey := []byte(viper.GetString("jwt.secret_key"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpexted singing method: %v\n", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		panic(err)
	}
	return token
}
