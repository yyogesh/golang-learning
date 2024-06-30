package file_handling

import (
	"fmt"
	"os"
)

func WriteFile(filename string, content string) error {
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

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func ReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(content), nil
}
