package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/models"
)

type ValidateIp2GeoRequestMiddleware struct {
	loggerImpl   common.Logger
	reqValidator *validator.Validate
}

func NewValidateIp2GeoRequestMiddleware(loggerImpl common.Logger) *ValidateIp2GeoRequestMiddleware {
	return &ValidateIp2GeoRequestMiddleware{
		loggerImpl:   loggerImpl,
		reqValidator: validator.New(),
	}
}

func (m *ValidateIp2GeoRequestMiddleware) ValidateIp2GeoRequest(ctx *gin.Context) {
	var req models.Ip2GeoRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		m.loggerImpl.Error("failed parsing json body - %v\n", err)
		return
	}

	err = m.reqValidator.Struct(req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		m.loggerImpl.Error("request missing data - %v\n", err)
		return
	}

	ctx.Set("ip2geo", req)
}
