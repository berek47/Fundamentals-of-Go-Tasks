package usecases

import (
	"errors"
)

type UserUsecase struct {
	UserRepo repositories.UserRepository
}

func (u *UserUsecase) Register(username, password, role string) (domain.User, error) {
	hashedPassword, err := infrastructure.HashPassword(password)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}
	return u.UserRepo.Register(user)
}

func (u *UserUsecase) Login(username, password string) (domain.User, error) {
	user, err := u.UserRepo.Authenticate(username, password)
	if err != nil {
		return domain.User{}, err
	}
	err = infrastructure.CheckPasswordHash(user.Password, password)
	if err != nil {
		return domain.User{}, errors.New("invalid credentials")
	}
	return user, nil
}