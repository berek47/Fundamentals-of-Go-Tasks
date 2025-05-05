package repositories

import (
	"errors"
)

type UserRepository interface {
	Register(user domain.User) (domain.User, error)
	Authenticate(username, password string) (domain.User, error)
}

type InMemoryUserRepo struct {
	users     []domain.User
	idCounter int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users:     []domain.User{},
		idCounter: 1,
	}
}

func (r *InMemoryUserRepo) Register(user domain.User) (domain.User, error) {
	for _, u := range r.users {
		if u.Username == user.Username {
			return domain.User{}, errors.New("username already exists")
		}
	}
	user.ID = r.idCounter
	r.idCounter++
	r.users = append(r.users, user)
	return user, nil
}

func (r *InMemoryUserRepo) Authenticate(username, _ string) (domain.User, error) {
	for _, u := range r.users {
		if u.Username == username {
			return u, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}