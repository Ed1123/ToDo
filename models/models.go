package models

import "time"

var idCounter = -1

type Task struct {
	Id            int
	Name          string
	Description   string
	Completed     bool
	DateCreated   time.Time
	DateUpdated   time.Time
	DateCompleted time.Time
	Archived      bool
}

func NewTask(name, description string) Task {
	idCounter++
	return Task{
		Id:          idCounter,
		Name:        name,
		Description: description,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}
}

var toDoIdCounter = -1

type ToDo struct {
	Id        int
	Tasks     []Task
	Date      time.Time
	Completed bool
	Archived  bool
}

func NewToDo() ToDo {
	toDoIdCounter++
	return ToDo{
		Id:   toDoIdCounter,
		Date: time.Now(),
	}
}
