package main

import (
	"resumme-builder/api"
	"resumme-builder/internal/utils/logger"
)

func main() {
	err := api.New().Run()
	if err != nil {
		logger.Log.Fatal(err)
	}
}
