package storage

import "TaskTracker/models"

var (
	tasks  = []models.Task{}
	nextId = 1
)

func GetAll() []models.Task {
	return tasks
}

func Add(title string) models.Task {
	task := models.Task{
		Id:        nextId,
		Title:     title,
		Completed: false,
	}

	nextId++
	tasks = append(tasks, task)
	return task
}
