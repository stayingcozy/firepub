package main

import (
	"fmt"
	"os"
)

func readUid(path_to_uid string) string {

	// Read the contents of the file
	content, err := os.ReadFile(path_to_uid)
	if err != nil {
		// Handle error
		fmt.Println("Error reading file:", err)
		return ""
	}

	// Convert the byte slice to a string
	uid_text := string(content)

	return uid_text
}