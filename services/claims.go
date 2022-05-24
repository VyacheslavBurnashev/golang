package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID string `json: "email"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	refreshTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	refreshClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims)
	tokenString, err := token.SignedString("")
	refreshTokenString, err := refreshToken.SignedString("")
	return tokenString, refreshTokenString, err
}

func GenerateNonAuthToken(userID string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString("")

	return tokenString, err
}

func DecodeNonAuthToken(tkStr string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tkStr, claims, func(t *jwt.Token) (interface{}, error) {
		return "", nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}
	if !token.Valid {
		return "", err
	}
	return claims.UserID, nil
}

func DecodeRefreshToken(tkStr string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tkStr, claims, func(t *jwt.Token) (interface{}, error) {
		return "", nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}
	if !token.Valid {
		return "", err
	}
	return claims.UserID, nil
}
