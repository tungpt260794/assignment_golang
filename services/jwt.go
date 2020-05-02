package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const hmacSecret = "iQasdq4MMdY2wxZCpAm1SxpbkGQopM4wx9QLgtVaHfjGCavuMLcuAZG6CvFxJaMd"

// TokenClaims ...
type TokenClaims struct {
	jwt.StandardClaims
	AccountID int `json:"account_id"`
}

// CreateToken ...
func CreateToken(accountID int) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		AccountID: accountID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Phan Thanh Tung - WD01 - VTA Academy",
			ExpiresAt: time.Now().Unix() + 24*60*60, // will be expired in 1 hour
			NotBefore: time.Now().Unix(),
		},
	})

	tokenStr, err = token.SignedString([]byte(hmacSecret))
	return
}

// VerifyToken ...
func VerifyToken(tokenStr string) (claims *TokenClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}

		return []byte(hmacSecret), nil
	})

	if err != nil {
		return
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Invalid token")
	}
}
