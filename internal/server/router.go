package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/controllers"
	"github.com/yairp7/ip2geo/internal/middlewares"
	"github.com/yairp7/ip2geo/internal/services"
	"github.com/yairp7/ip2geo/internal/services/geo"
)

var healthController *controllers.HealthController
var geoController *controllers.GeoController

var validateIp2GeoReqMiddleware *middlewares.ValidateIp2GeoRequestMiddleware
var ip2GeoCacheMiddleware *middlewares.Ip2GeoCacheMiddleware

func NewRouter(loggerImpl common.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthController = controllers.NewHealthController(loggerImpl)
	router.GET("/health", healthController.Status)

	cacher := services.NewInMemCache()

	geoController = controllers.NewGeoController(
		loggerImpl,
		geo.NewGeoService(
			loggerImpl,
			geo.WithIp2GeoHandler(geo.NewIpApiHandler(loggerImpl)),
			geo.WithIp2GeoHandler(geo.NewIpInfoHandler(loggerImpl)),
		),
		cacher,
	)

	validateIp2GeoReqMiddleware = middlewares.NewValidateIp2GeoRequestMiddleware(loggerImpl)
	ip2GeoCacheMiddleware = middlewares.NewIp2GeoCacheMiddleware(loggerImpl, cacher)
	router.POST(
		"/ip2geo",
		validateIp2GeoReqMiddleware.ValidateIp2GeoRequest,
		ip2GeoCacheMiddleware.GetIp2GeoFromCache,
		geoController.Ip2Geo,
	)

	return router
}

func ShutdownRouter() {
	fmt.Println("Router Shutdown")
	healthController.Close()
	geoController.Close()
}
