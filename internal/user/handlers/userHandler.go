package handlers

import (
	"encoding/json"
	"go_api/world/internal/user/domain"
	"net/http"
)

type UserHandler struct {
	userService domain.UserSvc
}

func NewUserHandler(userService domain.UserSvc) UserHandler {

	return UserHandler{userService: userService}
}

func (uh UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	data := &domain.User{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		panic(err)
	}

	success, _ := uh.userService.AddNewUser(data)

	var status int

	if success {
		status = http.StatusCreated
	} else {
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
}
