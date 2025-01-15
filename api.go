package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func (tasks *Tasks) addTask(Description string) (task, error) {
	if strings.TrimSpace(Description) == "" {
		return task{}, fmt.Errorf("task description cannot be empty")
	}

	if len(tasks.Tasks) >= 1000 {
		return task{}, fmt.Errorf("maximum task limit reached")
	}

	newTask := task{
		Id:          len(tasks.Tasks) + 1,
		Description: Description,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	tasks.Tasks = append(tasks.Tasks, newTask)
	return newTask, nil
}

func (tasks *Tasks) updateTask(id int, Description string) error {
	if strings.TrimSpace(Description) == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	found := false
	for i, t := range tasks.Tasks {
		if t.Id == id {
			tasks.Tasks[i].Description = Description
			tasks.Tasks[i].UpdateAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}
	return nil
}

func (tasks *Tasks) deleteTask(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid task ID: must be greater than 0")
	}

	var indexToDelete = -1
	for i, task := range tasks.Tasks {
		if task.Id == id {
			indexToDelete = i
			break
		}
	}
	if indexToDelete == -1 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	tasks.Tasks = append(tasks.Tasks[:indexToDelete], tasks.Tasks[indexToDelete+1:]...)
	return nil
}

func (tasks *Tasks) markDone(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid task ID: must be greater than 0")
	}

	found := false
	for i, t := range tasks.Tasks {
		if t.Id == id {
			if t.Status == "done" {
				return fmt.Errorf("task is already marked as done")
			}
			tasks.Tasks[i].Status = "done"
			tasks.Tasks[i].UpdateAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}
	return nil
}

func (tasks *Tasks) markProgress(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid task ID: must be greater than 0")
	}

	found := false
	for i, t := range tasks.Tasks {
		if t.Id == id {
			if t.Status == "in-progress" {
				return fmt.Errorf("task is already in progress")
			}
			tasks.Tasks[i].Status = "in-progress"
			tasks.Tasks[i].UpdateAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}
	return nil
}

func (tasks *Tasks) getTasks() ([]task, error) {
	if len(tasks.Tasks) == 0 {
		return nil, fmt.Errorf("no tasks found")
	}
	return tasks.Tasks, nil
}

func (tasks *Tasks) listDone() ([]task, error) {
	var doneTasks []task
	for _, t := range tasks.Tasks {
		if t.Status == "done" {
			doneTasks = append(doneTasks, t)
		}
	}
	if len(doneTasks) == 0 {
		return nil, fmt.Errorf("no completed tasks found")
	}
	return doneTasks, nil
}

func (tasks *Tasks) listInProgress() ([]task, error) {
	var inProgressTasks []task
	for _, t := range tasks.Tasks {
		if t.Status == "in-progress" {
			inProgressTasks = append(inProgressTasks, t)
		}
	}
	if len(inProgressTasks) == 0 {
		return nil, fmt.Errorf("no in-progress tasks found")
	}
	return inProgressTasks, nil
}

func (tasks *Tasks) listPending() ([]task, error) {
	var pendingTasks []task
	for _, t := range tasks.Tasks {
		if t.Status == "pending" {
			pendingTasks = append(pendingTasks, t)
		}
	}
	if len(pendingTasks) == 0 {
		return nil, fmt.Errorf("no pending tasks found")
	}
	return pendingTasks, nil
}

func saveTasks(filename string, tasks *Tasks) error {
	if tasks == nil {
		return fmt.Errorf("tasks cannot be nil")
	}

	if strings.TrimSpace(filename) == "" {
		return fmt.Errorf("filename cannot be empty")
	}

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("error marshaling tasks: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

func loadTasks(filename string) (*Tasks, error) {
	if strings.TrimSpace(filename) == "" {
		return nil, fmt.Errorf("filename cannot be empty")
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &Tasks{}, nil
		}
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var tasks Tasks
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling tasks: %v", err)
	}

	return &tasks, nil
}
