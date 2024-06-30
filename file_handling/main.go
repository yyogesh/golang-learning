package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the filename: ")
	input, _ := reader.ReadString('\n')
	filename := strings.TrimSpace(input)

	// Check if file exists
	if fileExists(filename) {
		fmt.Printf("File %s already exists.\n", filename)
	} else {
		fmt.Printf("File %s does not exist. It will be created.\n", filename)
	}

	// fmt.Scan
	fmt.Print("Enter text to write to file: ")
	inputContent, _ := reader.ReadString('\n')

	// Write input to file
	err := writeFile(filename, inputContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("File %s has been created/updated.\n", filename)

	// Read data from file
	content, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Data read from %s: %s\n", filename, content)
}

func writeFile(filename string, content string) error {
	// 4 for read permission (r)
	// 2 for write permission (w)
	// 1 for execute permission (x)
	// 0644
	// Owner can read and write: 6 (4+2)
	// Group can read: 4
	// Others can read: 4
	// 0755
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func readFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(content), nil
}
