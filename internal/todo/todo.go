// Package todo provides primitives for managing a persistent todo list,
// including data structures and file-based storage logic.
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Item struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

const filename = "todos.json"

func Load() ([]Item, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Item{}, nil
		}
		return nil, err
	}
	var todos []Item
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func Save(todos []Item) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0o644)
}

func Add(task string) (Item, error) {
	todos, _ := Load()
	newID := 1

	if len(todos) > 0 {
		newID = todos[len(todos)-1].ID + 1
	}

	newItem := Item{ID: newID, Task: task, Completed: false}
	todos = append(todos, newItem)
	err := Save(todos)
	return newItem, err
}

func Delete(id int) error {
	todos, err := Load()
	if err != nil {
		return err
	}

	newTodos := []Item{}
	found := false

	for _, item := range todos {
		if item.ID == id {
			found = true
			continue
		}
		newTodos = append(newTodos, item)
	}

	if !found {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	return Save(newTodos)
}

func Edit(id int, newTask string, completed bool) error {
	todos, err := Load()
	if err != nil {
		return err
	}

	found := false
	for i := range todos {
		if todos[i].ID == id {
			if newTask != "" {
				todos[i].Task = newTask
			}

			todos[i].Completed = completed
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	return Save(todos)
}
