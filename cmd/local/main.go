package main

import (
	"resumme-builder/internal/models"
	"resumme-builder/internal/pkg/pdf"
	"resumme-builder/internal/pkg/resume"
	"resumme-builder/internal/utils/logger"
)

func main() {
	logger.Log.Info("Generating output")

	resumeData, err := resume.ReadLocalData()
	if err != nil {
		logger.Log.Error(err)
		return
	}

	htmlFile, err := resume.ParseToHtml(resumeData)
	if err != nil {
		logger.Log.Fatal(err)
	}

	pdfData, err := pdf.GenerateFromHtml(htmlFile)
	if err != nil {
		logger.Log.Fatal(err)
	}

	if err := pdf.Write(models.OutputPdfFile, pdfData); err != nil {
		logger.Log.Fatal(err)
	}
}
