package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resumme-builder/internal/api/router"
	"resumme-builder/internal/utils/config"
	"resumme-builder/internal/utils/logger"
)

type Api struct {
	config config.ApiConfig
	router *gin.Engine
}

func New() *Api {
	api := &Api{}
	api.setup()
	return api
}

func (api *Api) setup() {
	api.config = config.LoadApiConfig()
	api.router = router.Init()
}

func (api *Api) Run() error {
	logger.Log.Info(fmt.Sprintf("%s API running on Port %d", api.config.AppName, api.config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", api.config.Port), api.router)
	if err != nil {
		return err
	}
	return nil
}