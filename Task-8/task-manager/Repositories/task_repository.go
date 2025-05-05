package repositories

import domain "task-manager/Domain"

type TaskRepository interface {
	CreateTask(task domain.Task) domain.Task
	GetAllTasks() []domain.Task
}