package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucasnevespereira/resb/configs"
	"github.com/lucasnevespereira/resb/internal/api/router"
	"github.com/lucasnevespereira/resb/internal/chrome"
	"github.com/lucasnevespereira/resb/internal/render"
	"github.com/lucasnevespereira/resb/internal/utils/logger"
	"net/http"
	"os"
)

type Api struct {
	config configs.ApiConfig
	router *gin.Engine
}

func New() (*Api, error) {
	renderer, err := render.New(os.DirFS("ui"), chrome.NewPrinter())
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
