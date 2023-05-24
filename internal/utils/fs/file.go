package fs

import (
	"github.com/pkg/errors"
	"os"
)

// EnsureNonEmptyFile ensures file exits and is non empty
func EnsureNonEmptyFile(filePath string) error {
	if filePath == "" {
		return errors.New("missing file path")
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return errors.Wrap(err, "failed to read file information")
	}

	if fileInfo.Size() == 0 {
		return errors.New("file is empty")
	}

	return nil
}

// WriteFile writes data to the specified destination.
func WriteFile(destination string, data []byte) error {
	if err := os.WriteFile(destination, data, os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to write file")
	}

	return nil
}

// CreateFile creates a file to the specified destination.
func CreateFile(destination string) (*os.File, error) {
	htmlOut, err := os.Create(destination)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file")
	}

	return htmlOut, nil
}

// ReadFile reads data from specified file.
func ReadFile(file string) ([]byte, error) {
	fileData, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}

	return fileData, nil
}
