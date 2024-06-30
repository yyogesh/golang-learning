package main

import (
	"bufio"
	"fmt"
	"os"
	"package_example/file_handling"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {

	// Print a simple message in cyan
	color.Cyan("Hello, Cyan World!")

	// Print a red message
	color.Red("This is a red message!")

	// Print a yellow message
	color.Yellow("This is a yellow message!")

	// Print a green message
	color.Green("This is a green message!")

	fmt.Println("********************************")

	str := "hello world"

	titleCase := cases.Title(language.English)

	fmt.Println(str, titleCase.String(str))

	reader := bufio.NewReader(os.Stdin)
	displayMessage()
	input, _ := reader.ReadString('\n')
	filename := strings.TrimSpace(input)

	// Check if file exists
	if file_handling.FileExists(filename) {
		fmt.Printf("File %s already exists.\n", filename)
	} else {
		fmt.Printf("File %s does not exist. It will be created.\n", filename)
	}

	// fmt.Scan
	fmt.Print("Enter text to write to file: ")
	inputContent, _ := reader.ReadString('\n')

	// Write input to file
	err := file_handling.WriteFile(filename, inputContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("File %s has been created/updated.\n", filename)

	// Read data from file
	content, err := file_handling.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Data read from %s: %s\n", filename, content)
}
