package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasnevespereira/resumme-builder/internal/api/handlers"
	"github.com/lucasnevespereira/resumme-builder/internal/render"
)

func Init(renderer *render.Renderer) *gin.Engine {
	router := gin.New()

	router.GET("/status", handlers.Status())
	router.POST("/pdf", handlers.GetPdf(renderer))

	return router
}
