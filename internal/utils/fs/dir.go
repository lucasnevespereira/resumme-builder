package fs

import (
	"fmt"
	"os"
)

// EnsureDir ensures directory exits
func EnsureDir(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return fmt.Errorf("EnsureDir - directory did not exist: %v", err)
	}
	return nil
}

// CreateDir ensures directory exits and creates it
func CreateDir(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, os.ModePerm)
		if err != nil {
			return fmt.Errorf("EnsureAndCreateDir - failed to create directory: %v", err)
		}
	}
	return nil
}
