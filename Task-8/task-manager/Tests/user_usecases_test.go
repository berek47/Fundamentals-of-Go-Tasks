package tests

import (
	domain "task-manager/Domain"
	"task-manager/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Register(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepo) Authenticate(username, password string) (domain.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	u := usecases.UserUsecase{UserRepo: mockRepo}
	user := domain.User{Username: "johndoe", Password: "pass", Role: "user"}

	mockRepo.On("Register", user).Return(user, nil)

	result, err := u.Register("johndoe", "pass", "user")
	assert.Nil(t, err)
	assert.Equal(t, "johndoe", result.Username)
	mockRepo.AssertExpectations(t)
}

func TestLogin_InvalidPassword(t *testing.T) {
	mockRepo := new(MockUserRepo)
	u := usecases.UserUsecase{UserRepo: mockRepo}
	storedUser := domain.User{Username: "johndoe", Password: "correct"}

	mockRepo.On("Authenticate", "johndoe", "wrong").Return(storedUser, nil)

	_, err := u.Login("johndoe", "wrong")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
}