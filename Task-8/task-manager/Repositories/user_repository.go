package repositories

import domain "task-manager/Domain"

type UserRepository interface {
	Register(user domain.User) (domain.User, error)
	Authenticate(username, password string) (domain.User, error)
}