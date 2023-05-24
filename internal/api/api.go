package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resumme-builder/configs"
	"resumme-builder/internal/api/router"
	"resumme-builder/internal/models"
	"resumme-builder/internal/pkg/parser"
	"resumme-builder/internal/pkg/pdf"
	"resumme-builder/internal/pkg/template"
	"resumme-builder/internal/services"
	"resumme-builder/internal/utils/logger"
)

type Api struct {
	config configs.ApiConfig
	router *gin.Engine
}

func New() *Api {
	api := &Api{}
	api.setup()
	return api
}

func (api *Api) setup() {
	templateManager := template.NewTemplateManager("ui")
	parser := parser.NewHTMLParser(models.OutputDir, models.OutputHtmlFile, templateManager)
	pdfGenerator := pdf.NewPDFGenerator()
	resumeService := services.NewResumeService(parser, pdfGenerator)
	api.config = configs.LoadApiConfig()
	api.router = router.Init(resumeService)
}

func (api *Api) Run() error {
	logger.Log.Info(fmt.Sprintf("%s API running on Port %d", api.config.AppName, api.config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", api.config.Port), api.router)
	if err != nil {
		return err
	}
	return nil
}
