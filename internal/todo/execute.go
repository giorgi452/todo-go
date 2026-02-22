package todo

import (
	"fmt"
	"strconv"
	"strings"
)

func Execute(args []string) {
	cmd := args[0]

	switch cmd {
	case "add":
		if len(args) < 2 {
			fmt.Println("❌ Error: Missing task description")
			return
		}
		task := strings.Join(args[1:], " ")
		item, _ := Add(task)
		fmt.Printf("✅ Added: %s (ID: %d)\n", item.Task, item.ID)

	case "list":
		items, _ := Load()
		fmt.Println("ID | Status | Task")
		for _, i := range items {
			status := " "
			if i.Completed {
				status = "x"
			}
			fmt.Printf("%d  |  [%s]   | %s\n", i.ID, status, i.Task)
		}

	case "edit":
		if len(args) < 3 {
			fmt.Println("Usage: edit <id> <new_task>")
			return
		}
		id, _ := strconv.Atoi(args[1])
		newTask := strings.Join(args[2:], " ")
		err := Edit(id, newTask, false)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("✏️ Item updated!")

	case "done":
		if len(args) < 2 {
			fmt.Println("Usage: done <id>")
			return
		}

		id, _ := strconv.Atoi(args[1])
		err := Edit(id, "", true)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("✅ Item marked as done!")

	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: delete <id>")
			return
		}
		for _, item := range args[1:] {
			id, _ := strconv.Atoi(item)
			err := Delete(id)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
				return
			}
			fmt.Printf("🗑️ Deleted item %d\n", id)
		}

	default:
		fmt.Println("Unknown command. Try: add, list, edit, done, delete, or exit")
	}
}
