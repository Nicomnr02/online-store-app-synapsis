package repositories

import (
	"online_app_store/model"

	"gorm.io/gorm"
)

type CredentialRepositoryInterface interface {
	CreateCredential(newCreds model.Credential) (model.Credential, error)
	GetCredentialByUserID(id int) (model.Credential, error)
	UpdateCredential(creds model.Credential) (model.Credential, error)
	DeleteCredential(creds model.Credential) error
}

type CredentialRepository struct {
	db *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) *CredentialRepository {
	return &CredentialRepository{db}

}

func (u *CredentialRepository) CreateCredential(newCreds model.Credential) (model.Credential, error) {

	if err := u.db.Create(&newCreds).Error; err != nil {
		return model.Credential{}, err
	}

	return newCreds, nil
}

func (u *CredentialRepository) GetCredentialByUserID(id int) (model.Credential, error) {
	var credsByUserID model.Credential
	if err := u.db.Where(&model.Credential{UserID: id}).First(&credsByUserID).Error; err != nil {
		return model.Credential{}, err
	}
	return credsByUserID, nil
}

func (u *CredentialRepository) UpdateCredential(creds model.Credential) (model.Credential, error) {

	if err := u.db.Model(&creds).Where("user_id = ?", creds.UserID).Updates(&creds).Error; err != nil {
		return model.Credential{}, err
	}

	return creds, nil

}

func (u *CredentialRepository) DeleteCredential(creds model.Credential) error {
	if err := u.db.Unscoped().Where("user_id = ?", creds.UserID).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
