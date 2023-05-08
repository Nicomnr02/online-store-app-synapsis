package services

import (
	"online_app_store/model"
	"online_app_store/repositories"
)

type UserServiceInterface interface {
	Register(user model.User) (model.User, error)
}

type UserService struct {
	UserRepo repositories.UserRepositoryInterface
}

func NewUserService(userRepo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{userRepo}
}

func (us *UserService) Register(user model.User) (model.User, error) {

	if existedUser, err := us.UserRepo.CreateUser(user); err != nil {
		return model.User{}, err
	} else {
		return existedUser, nil
	}

}
