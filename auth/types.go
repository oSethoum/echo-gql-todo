package auth

import "github.com/golang-jwt/jwt"

type Respose struct {
	Errors interface{} `json:"errors"`
	Data   interface{} `json:"data"`
}

type ContextKey struct {
	Key string
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}
