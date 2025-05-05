package usecases

import (
	"errors"
	domain "task-manager/Domain"
	repositories "task-manager/Repositories"
)

type UserUsecase struct {
	UserRepo repositories.UserRepository
}

func (u *UserUsecase) Register(username, password, role string) (domain.User, error) {
	if username == "" || password == "" {
		return domain.User{}, errors.New("missing required fields")
	}
	user := domain.User{
		Username: username,
		Password: password,
		Role:     role,
	}
	return u.UserRepo.Register(user)
}

func (u *UserUsecase) Login(username, password string) (domain.User, error) {
	user, err := u.UserRepo.Authenticate(username, password)
	if err != nil {
		return domain.User{}, err
	}
	if user.Password != password {
		return domain.User{}, errors.New("invalid credentials")
	}
	return user, nil
}
