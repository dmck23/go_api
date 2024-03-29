package handlers

import (
	"encoding/json"
	"go_api/world/internal/user/domain"
	"net/http"
)

type UserHandler struct {
	repo domain.UserDb
}

func NewUserHandler(repo domain.UserDb) UserHandler {

	return UserHandler{repo: repo}
}

func (uh UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	data := &domain.UserRequestModel{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		panic(err)
	}

	success, _ := uh.repo.CreateUser(data)

	var status int

	if success {
		status = http.StatusCreated
	} else {
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
}
