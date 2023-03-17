package entity

import "github.com/golang-jwt/jwt/v4"

type User struct {
	UserId   int    `json:"id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaim struct {
	UserId   int    `json:"id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}
