package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/models"
)

type Ip2GeoCacheMiddleware struct {
	loggerImpl common.Logger
	cacher     common.Cacher
}

func NewIp2GeoCacheMiddleware(loggerImpl common.Logger, cacher common.Cacher) *Ip2GeoCacheMiddleware {
	return &Ip2GeoCacheMiddleware{
		loggerImpl: loggerImpl,
		cacher:     cacher,
	}
}

func (m *Ip2GeoCacheMiddleware) GetIp2GeoFromCache(ctx *gin.Context) {
	req := ctx.Value("ip2geo").(models.Ip2GeoRequest)
	resp, err := m.cacher.Get(ctx, req.IP)
	if err == nil {
		handlerResponse := resp.(models.Ip2GeoHandlerResponse)
		response := models.Ip2GeoResponse{
			ReqID:                 req.ReqID,
			Ip2GeoHandlerResponse: handlerResponse,
		}
		ctx.AbortWithStatusJSON(http.StatusOK, response)
	}
}
