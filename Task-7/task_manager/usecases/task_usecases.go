package usecases

type TaskUsecase struct {
	TaskRepo repositories.TaskRepository
}

func (t *TaskUsecase) Create(task domain.Task) domain.Task {
	return t.TaskRepo.CreateTask(task)
}

func (t *TaskUsecase) GetAll() []domain.Task {
	return t.TaskRepo.GetAllTasks()
}