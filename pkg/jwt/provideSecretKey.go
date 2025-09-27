package jwt

import "os"

func ProvideSecretKey() []byte {
	return []byte(os.Getenv("APP_SECRET_KEY"))
}
