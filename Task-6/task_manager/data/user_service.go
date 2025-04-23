package data

import (
	"errors"
	"task_manager/models"

	"golang.org/x/crypto/bcrypt"
)

var users = []models.User{}
var idCounter = 1

func RegisterUser(username, password, role string) (models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return models.User{}, errors.New("username already exists")
		}
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := models.User{
		ID:       idCounter,
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}
	idCounter++
	users = append(users, newUser)
	return newUser, nil
}

func AuthenticateUser(username, password string) (models.User, error) {
	for _, user := range users {
		if user.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				return models.User{}, errors.New("invalid password")
			}
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}