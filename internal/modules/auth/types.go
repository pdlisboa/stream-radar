package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IAuthService interface {
	Login(req LoginRequest) (string, error)
}
