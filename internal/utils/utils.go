package utils

import (
	"log"
	"os"
)

func GetFilePathAsUrl(filename string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return "file://" + path + "/" + filename
}

func EnsureOutputDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Printf("EnsureOutputDir: %v", err)
		}
	}
}