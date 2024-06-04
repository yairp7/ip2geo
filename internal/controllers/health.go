package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yairp7/ip2geo/internal/common"
)

type HealthController struct {
	BaseController
}

func NewHealthController(loggerImpl common.Logger) *HealthController {
	return &HealthController{
		BaseController: NewBaseController("HealthController", loggerImpl),
	}
}

func (c *HealthController) Status(ctx *gin.Context) {
	if !c.isActive {
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	ctx.String(http.StatusOK, "OK!")
}
