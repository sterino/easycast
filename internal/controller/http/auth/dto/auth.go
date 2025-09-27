package dto

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar,omitempty"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
