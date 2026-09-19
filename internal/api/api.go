package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"resumme-builder/configs"
	"resumme-builder/internal/api/router"
	"resumme-builder/internal/pkg/pdf"
	"resumme-builder/internal/render"
	"resumme-builder/internal/utils/logger"
)

type Api struct {
	config configs.ApiConfig
	router *gin.Engine
}

func New() (*Api, error) {
	renderer, err := render.New(os.DirFS("ui"), pdf.NewPDFGenerator())
	if err != nil {
		return nil, err
	}
	return &Api{
		config: configs.LoadApiConfig(),
		router: router.Init(renderer),
	}, nil
}

func (api *Api) Run() error {
	logger.Log.Info(fmt.Sprintf("%s API running on Port %d", api.config.AppName, api.config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", api.config.Port), api.router)
	if err != nil {
		return err
	}
	return nil
}
