package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resumme-builder/internal/models"
	"resumme-builder/internal/render"
	"resumme-builder/internal/utils/logger"
)

func Status() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

func GetPdf(renderer *render.Renderer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resumeData models.Resume

		if err := c.ShouldBindJSON(&resumeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "ShouldBindJSON : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		pdfData, err := renderer.PDF(c.Request.Context(), resumeData)
		if err != nil {
			logger.Log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  err.Error(),
				"status": http.StatusInternalServerError,
			})
			return
		}

		c.Data(http.StatusOK, "application/pdf", pdfData)
	}
}
