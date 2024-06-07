package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Note struct {
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Create new note and press enter:")
  fmt.Println(os.Args)

	if scanner.Scan() {
		input := scanner.Text()
		fmt.Println("New note:", input)
		appendToFile(input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file", err)
	}
}

func appendToFile(content string) {
	filePath := "notes.json"
	var notes []Note

	if _, err := os.Stat(filePath); err == nil {
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		if err := json.Unmarshal(fileData, &notes); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
	}

	newNote := Note{
		Content: content,
		ID:      len(notes) + 1,
	}
	notes = append(notes, newNote)

	updatedData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	if err := os.WriteFile(filePath, updatedData, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Note added successfully")
}
