package main

import (
	"log"
	"resume-builder/internal/pdf"
	"resume-builder/internal/resume"
	"resume-builder/utils"
)

func main() {
	err := resume.ParseToHtml(utils.BasicTemplateFiles)
	if err != nil {
		log.Fatal(err)
	}

	err = pdf.Generate()
	if err != nil {
		log.Fatal(err)
	}
}
