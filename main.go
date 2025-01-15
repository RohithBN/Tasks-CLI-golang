package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

func main() {
	filename := "tasks.json"
	tasks, err := loadTasks(filename)
	if err != nil {
		log.Fatal(err)
	}

	add := flag.String("add", "", "Add a task")
	list := flag.Bool("list", false, "List tasks")
	markdone := flag.Int("mark-done", 0, "Mark a task as done")
	markinprogress := flag.Int("mark-in-progress", 0, "Mark a task as in-progress")
	delete := flag.Int("delete", 0, "Delete a task")
	listDone := flag.Bool("list-done", false, "List done tasks")
	listinprogress := flag.Bool("list-in-progress", false, "List in-progress tasks")
	listPending := flag.Bool("list-pending", false, "List pending tasks")
	updateID := flag.Int("update-id", 0, "ID of the task to update")
	updateDescription := flag.String("update-description", "", "New description for the task")
	flag.Parse()

	if *add != "" {
		tasks.addTask(*add)
		fmt.Println("Task added")
		err = saveTasks(filename, tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *updateID != 0 && *updateDescription != "" {
		err = tasks.updateTask(*updateID, *updateDescription)
		if err != nil {
			log.Fatal(err)
		}
		err = saveTasks("tasks.json", tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *markdone != 0 {
		tasks.markDone(*markdone)
		fmt.Println("Task marked as done")
		err = saveTasks(filename, tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *markinprogress != 0 {
		tasks.markProgress(*markinprogress)
		fmt.Println("Task marked as in-progress")
		err = saveTasks(filename, tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *delete != 0 {
		tasks.deleteTask(*delete)
		fmt.Println("Task deleted")
		err = saveTasks(filename, tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
	if *list {
		tasks, err := tasks.getTasks()
		if err != nil {
			log.Fatal(err)
		}
		tasksJSON, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(tasksJSON))

	}
	if *listDone {
		tasks, err := tasks.listDone()
		if err != nil {
			log.Fatal(err)
		}
		tasksJSON, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(tasksJSON))
	}
	if *listinprogress {
		tasks, err := tasks.listInProgress()
		if err != nil {
			log.Fatal(err)
		}
		tasksJSON, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(tasksJSON))
	}
	if *listPending {
		tasks, err := tasks.listPending()
		if err != nil {
			log.Fatal(err)
		}
		tasksJSON, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(tasksJSON))
	}

}
