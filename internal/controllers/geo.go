package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/models"
	"github.com/yairp7/ip2geo/internal/services/geo"
)

type Ip2GeoResponse struct {
	ReqID string `json:"reqId"`
	models.Ip2GeoHandlerResponse
}

type GeoController struct {
	BaseController
	geoService *geo.GeoService
	cacher     common.Cacher
}

func NewGeoController(
	loggerImpl common.Logger,
	geoService *geo.GeoService,
	cacher common.Cacher,
) *GeoController {
	c := &GeoController{
		BaseController: NewBaseController("GeoController", loggerImpl),
		geoService:     geoService,
		cacher:         cacher,
	}

	c.RegisterService(c.geoService)

	return c
}

func (c *GeoController) Ip2Geo(ctx *gin.Context) {
	if !c.isActive {
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}

	c.RegisterOp()
	defer c.UnregisterOp()

	req := ctx.Value("ip2geo").(models.Ip2GeoRequest)

	serviceResponse, err := c.geoService.Ip2Geo(req.IP)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		c.loggerImpl.Error("failed resolving geo for ip %s - %v\n", req.IP, err)
		return
	}

	if c.cacher != nil {
		c.cacher.Set(ctx, req.IP, serviceResponse)
	}

	response := Ip2GeoResponse{
		ReqID:                 req.ReqID,
		Ip2GeoHandlerResponse: serviceResponse,
	}

	ctx.JSON(http.StatusOK, response)
}
