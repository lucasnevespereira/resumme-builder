package main

import (
	"log"
	"resume-builder/resume"
)

func main() {
	err := resume.ParseToHtml()
	if err != nil {
		log.Fatal(err)
	}
}
