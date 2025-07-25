package auth

import (
	"errors"
	"stream-radar/domain/model"
	"stream-radar/internal/database"
	"stream-radar/internal/utils"
)

type AuthService struct{}

var service *AuthService

func init() {
	service = &AuthService{}
}

func (svc AuthService) Login(req LoginRequest) (string, error) {
	var user model.User
	res := database.DB.Where("email = ?", req.Email).First(&user)
	if res.Error != nil {
		return "", errors.New("incorrect email or password")
	}

	if !utils.ComparePassword(user.PasswordHash, req.Password) {
		return "", errors.New("incorrect email or password")
	}

	token, err := utils.GenerateToken(user.Id)

	if err != nil {
		return "", err
	}
	return token, nil
}
