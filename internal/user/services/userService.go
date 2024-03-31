package services

import (
	"go_api/world/internal/user/domain"
	"go_api/world/internal/user/errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserDb
}

func NewUserService(repo domain.UserDb) domain.UserSvc {
	return UserService{repo: repo}
}

func (u UserService) AddNewUser(user *domain.User) (bool, error) {

	_, userNameErr := u.repo.FindUserByUserName(user.Username)

	if userNameErr == nil {
		log.Println(errors.UserAlreadyExists{})
		return false, errors.UserAlreadyExists{}
	}

	_, emailErr := u.repo.FindUserByEmail(user.Email)

	if emailErr == nil {
		log.Println(errors.EmailAlreadyInUse{})
		return false, errors.EmailAlreadyInUse{}
	}

	encryptedPassword, err := encryptPassword(user.Password)

	if err != nil {
		return false, err
	}

	user.Password = encryptedPassword

	saveData, _ := formatUser(user)

	result, err := u.repo.CreateUser(saveData)

	if err != nil {
		return false, err
	}

	log.Println(result)

	return true, nil
}

func encryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func formatUser(user *domain.User) ([]byte, error) {

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	user.UpdatedAt = time.Now()

	type my domain.User
	return bson.Marshal((*my)(user))

}
