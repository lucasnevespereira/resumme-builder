package router

import (
	"github.com/gin-gonic/gin"
	"resume-builder/api/handlers"
)

func Init() *gin.Engine {
	router := gin.New()

	router.GET("/status", handlers.Status())
	router.POST("/pdf", handlers.GetPdf())

	return router
}

