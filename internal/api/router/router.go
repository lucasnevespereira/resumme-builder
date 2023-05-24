package router

import (
	"github.com/gin-gonic/gin"
	"resumme-builder/internal/api/handlers"
	"resumme-builder/internal/services"
)

func Init(service *services.ResumeService) *gin.Engine {
	router := gin.New()

	router.GET("/status", handlers.Status())
	router.POST("/pdf", handlers.GetPdf(service))

	return router
}
