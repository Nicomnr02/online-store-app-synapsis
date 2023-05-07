package repositories

import (
	"online_app_store/model"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(User model.User) (model.User, error)
	GetUserByID(id int) (model.User, error)
	UpdateUser(User model.User) (model.User, error)
	DeleteUser(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}

}

func (u *UserRepository) CreateUser(User model.User) (model.User, error) {
	if err := u.db.Create(&User).Error; err != nil {
		return model.User{}, err
	}

	return User, nil
}

func (u *UserRepository) GetUserByID(id int) (model.User, error) {
	var UserByID model.User
	if err := u.db.First(&UserByID, id).Error; err != nil {
		return model.User{}, err
	}
	return UserByID, nil
}

func (u *UserRepository) UpdateUser(User model.User) (model.User, error) {

	if err := u.db.Model(&User).Updates(&User).Error; err != nil {
		return model.User{}, err
	}

	return User, nil

}

func (u *UserRepository) DeleteUser(id int) error {
	if err := u.db.Unscoped().Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
