package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
  "time"
)

type Note struct {
	Content string `json:"content"`
	ID      int    `json:"id"`
  Timestamp time.Time `json: "timestamp"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
  if len(os.Args) >= 2 {
    if os.Args[1] == "new" {
    fmt.Println("New note:")
	if scanner.Scan() {
		input := scanner.Text()
		appendToFile(input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file", err)
	}
}
if os.Args[1] == "all" {
  fmt.Println("read all notes")
  readNotesFromFile()
}
}
if len(os.Args) < 2 {
  fmt.Println("Type 'new' to create a new note or 'all' to see all your notes.")
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
    Timestamp: time.Now(),
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

func readNotesFromFile() {
	filePath := "notes.json"
	if _, err := os.Stat(filePath); err == nil {
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		var notes []Note
		if err := json.Unmarshal(fileData, &notes); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		for _, note := range notes {
			fmt.Printf("Note ID: %d\nContent: %s\nTimestamp: %s\n\n", note.ID, note.Content, note.Timestamp.Format("02-01-2006 15:04"))
		}
	} else {
		fmt.Println("No notes found.")
	}
}
