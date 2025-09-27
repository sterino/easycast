package jwt

import "github.com/dgrijalva/jwt-go"

type Profile struct {
	PhoneNumber string `json:"phone_number"`
}

type Claims struct {
	Profile Profile
	jwt.StandardClaims
}
