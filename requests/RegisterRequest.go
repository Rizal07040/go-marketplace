package requests

type RegisterRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"PasswordConfirm"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
