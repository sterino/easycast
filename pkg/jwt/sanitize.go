package jwt

// Sanitize makes sure that the given token is valid.
func Sanitize(token string) bool {
	header, payload, signature := Split(token)
	return header != nil && payload != nil && signature != nil
}
