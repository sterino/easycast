package password

import "golang.org/x/crypto/bcrypt"

func Generate(password string) (string, error) {
	bslice := []byte(password)

	generated, err := bcrypt.GenerateFromPassword(bslice, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashed := string(generated)

	return hashed, nil
}
