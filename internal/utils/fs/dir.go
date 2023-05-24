package fs

import (
	"os"
	"resumme-builder/internal/utils/logger"
)

// EnsureDir ensures directory exits and creates it
func EnsureDir(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, os.ModePerm)
		if err != nil {
			logger.Log.Errorf("EnsureDir - failed to create directory: %v", err)
		}
	}
}
