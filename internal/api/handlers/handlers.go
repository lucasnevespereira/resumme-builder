package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resumme-builder/internal/models"
	"resumme-builder/internal/pkg/pdf"
	"resumme-builder/internal/pkg/resume"
	"resumme-builder/internal/utils/logger"
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

		c.Writer.Header().Set("Content-type", "application/pdf")
		c.File(models.OutputPdfFile)
	}
}
