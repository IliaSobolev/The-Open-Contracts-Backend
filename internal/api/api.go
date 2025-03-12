package api

import (
	"TOC/internal/api/codeBlock"
	"TOC/internal/api/middleware"
	"TOC/pkg/domain"
	"TOC/pkg/utils"
	"github.com/gin-gonic/gin"
)

type API struct {
	r *gin.Engine

	cb *codeBlock.Module
}

func NewAPI(r *gin.Engine, cb domain.CodeBlockUsecase) *API {
	api := &API{r: r}

	api.r.Use(middleware.Logger)

	api.cb = codeBlock.NewModule(r, cb)

	return api
}

func (api *API) Start() error {
	err := api.r.Run(":" + utils.GetEnv("PORT", "1610"))
	if err != nil {
		return err
	}
	return nil
}
