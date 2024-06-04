package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/models"
)

func GetIp2GeoFromCache(loggerImpl common.Logger, cacher common.Cacher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Value("ip2geo").(models.Ip2GeoRequest)
		resp, err := cacher.Get(ctx, req.IP)
		if err != nil {
			return
		}

		handlerResponse := resp.(models.Ip2GeoHandlerResponse)
		response := models.Ip2GeoResponse{
			ReqID:                 req.ReqID,
			Ip2GeoHandlerResponse: handlerResponse,
		}
		ctx.AbortWithStatusJSON(http.StatusOK, response)
	}
}
