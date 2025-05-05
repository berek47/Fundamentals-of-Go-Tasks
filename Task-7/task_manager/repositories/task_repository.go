package repositories

type TaskRepository interface {
	CreateTask(task domain.Task) domain.Task
	GetAllTasks() []domain.Task
}

type InMemoryTaskRepo struct {
	tasks     []domain.Task
	idCounter int
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{
		tasks:     []domain.Task{},
		idCounter: 1,
	}
}

func (r *InMemoryTaskRepo) CreateTask(task domain.Task) domain.Task {
	task.ID = r.idCounter
	r.idCounter++
	r.tasks = append(r.tasks, task)
	return task
}

func (r *InMemoryTaskRepo) GetAllTasks() []domain.Task {
	return r.tasks
}