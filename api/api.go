package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resume-builder/api/router"
	"resume-builder/utils"
)

type Api struct {
	config utils.Config
	router *gin.Engine
}

func New() *Api {
	api := &Api{}
	api.setup()
	return api
}

func (api *Api) setup() {
	api.config = utils.LoadConfig()
	api.router = router.Init()
}

func (api *Api) Run() error {
	utils.Logger.Info(fmt.Sprintf("%s API running on Port %d", api.config.AppName, api.config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", api.config.Port), api.router)
	if err != nil {
		return err
	}
	return nil
}
