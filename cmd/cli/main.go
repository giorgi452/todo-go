package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo-go/internal/todo"
)

func main() {
	if len(os.Args) > 1 {
		todo.Execute(os.Args[1:])
		return
	}

	fmt.Println("--- Welcome to Todo ---")
	fmt.Println("Enter command (add, list, edit, done, delete, help) or 'exit' to quit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		if args[0] == "exit" || args[0] == "quit" {
			break
		}

		todo.Execute(args)
	}
}
