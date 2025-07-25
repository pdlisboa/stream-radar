package user

import "stream-radar/domain/model"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type FindUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type IUserService interface {
	Create(req CreateUserRequest) (model.User, error)
	Find(req FindUserRequest) (model.User, error)
}
