package security

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Jwt = string

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  Jwt `json:"access_token" xml:"access_token"`
	RefreshToken Jwt `json:"refresh_token" xml:"refresh_token"`
}

func GenerateJwt(secret, subject string, userId uint, expiresInMinutes int64) (Jwt, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expiresInMinutes))),
		},
	})
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GenerateJwtPair(secret, subject string, userId uint, accessTokenExpiresInMinutes, refreshTokenExpiresInMinutes int64) (TokenPair, error) {
	accessToken, err := GenerateJwt(secret, subject, userId, accessTokenExpiresInMinutes)
	if err != nil {
		return TokenPair{}, err
	}
	refreshToken, err := GenerateJwt(secret, subject, userId, refreshTokenExpiresInMinutes)
	if err != nil {
		return TokenPair{}, err
	}
	return TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
