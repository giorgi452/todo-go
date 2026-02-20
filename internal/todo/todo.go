package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Item struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

const filename = "todos.json"

// Load reads todos from the JSON file
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
	return os.WriteFile(filename, data, 0644)
}

func Add(task string) (Item, error) {
	todos, _ := Load()
	newID := 1

	if len(todos) > 0 {
		newID = todos[len(todos)-1].ID + 1
	}

	newItem := Item{ID: newID, Task: task, Done: false}
	todos = append(todos, newItem)
	err := Save(todos)
	return newItem, err
}
