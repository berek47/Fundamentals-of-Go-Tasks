package tests

import (
	domain "task-manager/Domain"
	"task-manager/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepo struct {
	mock.Mock
}

func (m *MockTaskRepo) CreateTask(task domain.Task) domain.Task {
	args := m.Called(task)
	return args.Get(0).(domain.Task)
}

func (m *MockTaskRepo) GetAllTasks() []domain.Task {
	args := m.Called()
	return args.Get(0).([]domain.Task)
}

func TestCreateTask(t *testing.T) {
	repo := new(MockTaskRepo)
	uc := usecases.TaskUsecase{TaskRepo: repo}
	task := domain.Task{Title: "Write Tests", User: "jane"}

	repo.On("CreateTask", task).Return(task)
	result := uc.Create(task)
	assert.Equal(t, "Write Tests", result.Title)
	repo.AssertExpectations(t)
}

func TestGetAllTasks(t *testing.T) {
	repo := new(MockTaskRepo)
	uc := usecases.TaskUsecase{TaskRepo: repo}
	tasks := []domain.Task{{Title: "Task1"}, {Title: "Task2"}}

	repo.On("GetAllTasks").Return(tasks)
	result := uc.GetAll()
	assert.Len(t, result, 2)
	repo.AssertExpectations(t)
}