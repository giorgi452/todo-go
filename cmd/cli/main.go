package main

import (
	"fmt"
	"os"
	"strconv"

	"todo-go/internal/todo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo-cli [add|list] <content>")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		task := os.Args[2]
		item, _ := todo.Add(task)
		fmt.Printf("✅ Added: %s (ID: %d)\n", item.Task, item.ID)

	case "list":
		items, _ := todo.Load()
		fmt.Println("ID | Status | Task")
		for _, i := range items {
			status := " "
			if i.Completed {
				status = "x"
			}
			fmt.Printf("%d  |  [%s]   | %s\n", i.ID, status, i.Task)
		}

	case "edit":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo-cli edit <id> <new_task>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		newTask := os.Args[3]

		err := todo.Edit(id, newTask, false)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("✏️ Item updated!")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide an ID to delete.")
			return
		}
		idStr := os.Args[2]

		id, _ := strconv.Atoi(idStr)

		err := todo.Delete(id)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}
		fmt.Printf("🗑️ Deleted item %d\n", id)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])

		err := todo.Edit(id, "", true)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("✅ Item marked as done!")
	}
}
