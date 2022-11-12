package main

import (
	"resumme-builder/api"
	"resumme-builder/utils"
)

func main() {
	err := api.New().Run()
	if err != nil {
		utils.Logger.Fatal(err)
	}
}


