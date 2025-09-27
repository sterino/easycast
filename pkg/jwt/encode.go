package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// issuer is used to set an issuer to the json web tokens.
const issuer = "marcos.huck.com.ar"

// expirationTime is the maximum duration that a JWT should last.
const expirationTime = time.Minute * 10

func Encode(body JWT, secretKey []byte) (*string, *int64, error) {
	expiresAt := time.Now().Add(expirationTime).Unix()

	claims := Claims{
		Profile: Profile{
			PhoneNumber: body.PhoneNumber,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "marketplace",
			Subject:   body.UUID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString(secretKey)
	if err != nil {
		return nil, nil, err
	}

	return &result, &expiresAt, nil
}
