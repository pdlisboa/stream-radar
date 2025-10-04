package user

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"stream-radar/domain/model"
	"stream-radar/internal/database"
	"stream-radar/internal/logger"
	"stream-radar/internal/utils"
	"time"
)

type UserService struct {
}

var service *UserService

func init() {
	service = &UserService{}
}

func (svc UserService) Create(req CreateUserRequest) (model.User, error) {
	log := logger.GetInstance()
	user := model.User{
		Email:        req.Email,
		Name:         req.Name,
		PasswordHash: utils.GeneratePassword(req.Password),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	res := database.DB.Create(&user)

	if res.Error != nil {
		log.Error(fmt.Sprintf("Error creating user %s", req.Name), zap.Error(res.Error))
		return model.User{}, res.Error
	}

	return user, nil
}

func (svc UserService) Find(req FindUserRequest) (model.User, error) {
	log := logger.GetInstance()

	var user model.User
	res := database.DB.Where("email = ? or name = ?", req.Email, req.Name).First(&user)

	if res.Error != nil {
		log.Debug("User not found", zap.String("email", req.Email))
		return model.User{}, errors.New("user not found")
	}

	return user, nil

}

func (svc UserService) Get(id uint) (model.User, error) {
	log := logger.GetInstance()

	var user model.User
	res := database.DB.Where("id = ?", id).First(&user)

	if res.Error != nil {
		log.Debug("User not found", zap.Uint("id", id))
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}
