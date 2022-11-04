package main

import (
	"resume-builder/api"
	"resume-builder/utils"
)

func main() {
	err := api.New().Run()
	if err != nil {
		utils.Logger.Fatal(err)
	}
}
