package api

import (
	"encoding/json"
	"net/http"
	"online_app_store/model"
	"online_app_store/services"
	"time"
)

type CredentialAPIinterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type CredentialAPI struct {
	credentialService services.CredentialServiceInterface
}

func NewCredentialAPI(credentialService services.CredentialServiceInterface) *CredentialAPI {
	return &CredentialAPI{credentialService}
}

func (ca *CredentialAPI) Login(w http.ResponseWriter, r *http.Request) {
	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	creds, err := ca.credentialService.Login(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	cookie := &http.Cookie{
		Name:    "synapsis",
		Value:   creds.SessionToken,
		Path:    "/",
		Expires: time.Now().Add(2 * time.Minute),
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.Response{Message: "Login Success", Data: data})

}

// func (ca *CredentialAPI) Logout(w http.ResponseWriter, r *http.Request) {
// 	var data model.User

// 	creds, err := ca.credentialService.Login(data)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(&model.Response{Message: "Login Success", Data: creds})

// }
