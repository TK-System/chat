package main

import (
	"chat/backend/infrastructures/cmd"
	"fmt"
)

func main() {
	err := cmd.ChatCmd.Execute()
	fmt.Printf("HELLO server starts\n")
	if err != nil {
		fmt.Printf("error: %w", err)
	}
	return
}
