package model

import "github.com/golang-jwt/jwt/v4"

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

type CustomClaims struct {
	Token string
	Ip    string
	jwt.RegisteredClaims
}
