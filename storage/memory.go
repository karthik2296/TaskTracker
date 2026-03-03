package storage

import "TaskTracker/models"

var (
	tasks  = []models.Task{}
	nextId = 1
)

func GetAll() []models.Task {
	return tasks
}

func NewAdd(title string) models.Task {
	task := models.Task{
		Id:        nextId,
		Title:     title,
		Completed: false,
	}

	nextId++
	tasks = append(tasks, task)
	return task
}

func TaskCompleted(title string) models.Task {
	for i, task := range tasks {
		if task.Title == title {
			tasks[i].Completed = true
			return tasks[i]
		}

	}
	return models.Task{}
}
