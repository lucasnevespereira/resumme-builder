package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resumme-builder/internal/models"
	"resumme-builder/internal/services"
	"resumme-builder/internal/utils/fs"
	"resumme-builder/internal/utils/logger"
)

func Status() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

func GetPdf(service *services.ResumeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resumeData models.Resume

		if err := c.ShouldBindJSON(&resumeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "ShouldBindJSON : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		htmlFile, err := service.Parser.ParseToHtml(resumeData)
		if err != nil {
			logger.Log.Fatal(err)
		}

		pdfData, err := service.Pdf.GenerateFromHTML(htmlFile)
		if err != nil {
			logger.Log.Fatal(err)
		}

		if err := fs.WriteFile(models.OutputPdfFile, pdfData); err != nil {
			logger.Log.Fatal(err)
		}

		c.Writer.Header().Set("Content-type", "application/pdf")
		c.File(models.OutputPdfFile)
	}
}
