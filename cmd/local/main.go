package main

import (
	"encoding/json"
	"os"
	"resumme-builder/internal/models"
	"resumme-builder/internal/pdf"
	"resumme-builder/internal/resume"
	"resumme-builder/internal/utils"
)

func main() {
	utils.Logger.Info("Generating output")
	fileData, err := os.ReadFile(utils.ResumeDataFile)
	if err != nil {
		utils.Logger.Error(err)
	}

	var resumeData models.Resume
	err = json.Unmarshal(fileData, &resumeData)
	if err != nil {
		utils.Logger.Error(err)
	}

	htmlFile, err := resume.ParseToHtml(resumeData)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	htmlFileUrl := utils.GetFilePathAsUrl(htmlFile)

	pdfData, err := pdf.Generate(htmlFileUrl)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	if err := os.WriteFile(utils.OutputPdfFile, pdfData, 0644); err != nil {
		utils.Logger.Fatal(err)
	}
}
