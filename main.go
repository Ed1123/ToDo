package main

import (
	"fmt"

	"github.com/Ed1123/todo/models"
)

func main() {
	tasks := []models.Task{
		models.NewTask("Take out trash", ""),
		models.NewTask("Clean room", ""),
		models.NewTask("Do laundry", ""),
	}

	todaysTodo := models.NewToDo()
	todaysTodo.Tasks = tasks
	fmt.Println(todaysTodo)
}
