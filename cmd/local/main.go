package main

import (
	"encoding/json"
	"os"
	"resume-builder/internal/pdf"
	"resume-builder/internal/resume"
	"resume-builder/shared/models"
	"resume-builder/utils"
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
