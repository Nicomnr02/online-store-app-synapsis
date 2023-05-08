package api

import (
	"encoding/json"
	"net/http"
	"online_app_store/model"
	"online_app_store/services"
)

type UserAPIINterface interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type UserAPI struct {
	userService services.UserServiceInterface
}

func NewUserAPI(userService services.UserServiceInterface) *UserAPI {
	return &UserAPI{userService}
}

func (ua *UserAPI) Register(w http.ResponseWriter, r *http.Request) {
	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	newAcc, err := ua.userService.Register(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&model.Response{Message: "Register Success", Data: newAcc})

}
