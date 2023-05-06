package services

import (
	"errors"
	"online_app_store/model"
	"online_app_store/repositories"
	"online_app_store/vars"

	"github.com/golang-jwt/jwt/v5"
)

type CredentialServiceInterface interface {
	Login(acc model.User) (model.Credential, error)
	Logout(acc model.User) (model.Credential, error)
}

type CredentialService struct {
	userRepo  repositories.UserRepositoryInterface
	credsRepo repositories.CredentialRepositoryInterface
}

func NewCredentialService(userRepo repositories.UserRepositoryInterface, credentialRepo repositories.CredentialRepositoryInterface) CredentialServiceInterface {
	return &CredentialService{userRepo, credentialRepo}
}

func (cs *CredentialService) Login(acc model.User) (model.Credential, error) {
	existedAcc, err := cs.userRepo.GetUserByID(acc.ID)
	if err != nil { //creds not found
		return model.Credential{}, err
	}

	if acc.Username != existedAcc.Username || acc.Password != existedAcc.Password {
		return model.Credential{}, errors.New("wrong username or password")
	}

	jwtToken, err := generateJWT(acc.ID)
	if err != nil {
		return model.Credential{}, err
	}

	creds, err := cs.credsRepo.CreateCredential(model.Credential{UserID: acc.ID, SessionToken: jwtToken})
	if err != nil {
		return model.Credential{}, err
	}

	return creds, nil

}

func (cs *CredentialService) Logout(acc model.User) (model.Credential, error) {
	creds, err := cs.credsRepo.GetCredentialByUserID(acc.ID)
	if err != nil {
		return model.Credential{}, err
	}

	err = cs.credsRepo.DeleteCredential(creds)
	if err != nil {
		return model.Credential{}, err
	}

	deletedCreds := creds
	return deletedCreds, nil

}

func generateJWT(userID int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	tokenString, err := token.SignedString(vars.JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
