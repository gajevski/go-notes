package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println("Error %w", err)
		return
	}
	bytes, err := file.WriteString("Test string 2")

	if err != nil {
		fmt.Println("Error %w", err)
		file.Close()
		return
	}

	fmt.Println(bytes, "File saved successfully")

}
