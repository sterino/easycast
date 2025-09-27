package password

import "golang.org/x/crypto/bcrypt"

func Compare(plain, hashed string) bool {
	bPlain := []byte(plain)
	bHashed := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(bHashed, bPlain)
	if err != nil {
		return false
	}
	return true
}
