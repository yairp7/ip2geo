package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/models"
)

func ValidateIp2GeoRequest(loggerImpl common.Logger) gin.HandlerFunc {
	reqValidator := validator.New()
	return func(ctx *gin.Context) {
		var req models.Ip2GeoRequest
		err := ctx.Bind(&req)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			loggerImpl.Error("failed parsing json body - %v\n", err)
			return
		}

		err = reqValidator.Struct(req)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			loggerImpl.Error("request missing data - %v\n", err)
			return
		}

		ctx.Set("ip2geo", req)
	}
}
