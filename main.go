package main

import (
	"fmt"
	"net/http"

	"github.com/Ed1123/todo/components"
	"github.com/Ed1123/todo/models"
	"github.com/a-h/templ"
)

func main() {
	tasks := []models.Task{
		models.NewTask("Take out trash", ""),
		models.NewTask("Clean room", ""),
		models.NewTask("Do laundry", ""),
	}

	todaysTodo := models.NewToDo()
	todaysTodo.Tasks = tasks

	http.Handle("/", templ.Handler(components.Page([]models.ToDo{todaysTodo, todaysTodo})))

	port := ":3000"
	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}
