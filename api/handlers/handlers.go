package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"resumme-builder/internal/models"
	"resumme-builder/internal/pdf"
	"resumme-builder/internal/resume"
	"resumme-builder/internal/utils"
)

func Status() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

func GetPdf() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resumeData models.Resume

		if err := c.ShouldBindJSON(&resumeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "ShouldBindJSON : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		htmlFileName, err := resume.ParseToHtml(resumeData)
		if err != nil {
			utils.Logger.Fatal(err)
		}

		htmlFileUrl := utils.GetFilePathAsUrl(htmlFileName)

		pdfData, err := pdf.Generate(htmlFileUrl)
		if err != nil {
			utils.Logger.Fatal(err)
		}

		if err := os.WriteFile(utils.OutputPdfFile, pdfData, 0644); err != nil {
			utils.Logger.Fatal(err)
		}

		c.Writer.Header().Set("Content-type", "application/pdf")
		c.File(utils.OutputPdfFile)
	}
}
