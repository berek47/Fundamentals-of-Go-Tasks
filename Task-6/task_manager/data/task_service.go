package data

import "task_manager/models"

var tasks = []models.Task{}
var taskIDCounter = 1

func CreateTask(task models.Task) models.Task {
	task.ID = taskIDCounter
	taskIDCounter++
	tasks = append(tasks, task)
	return task
}