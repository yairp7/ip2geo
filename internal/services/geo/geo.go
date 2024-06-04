package geo

import (
	"errors"

	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/models"
)

type Ip2GeoHandler interface {
	Name() string
	Ip2Geo(ip string) (models.Ip2GeoHandlerResponse, error)
}

type GeoServiceOption func(*GeoService)

type GeoService struct {
	loggerImpl common.Logger
	handlers   []Ip2GeoHandler
}

var NoHandlersError = errors.New("no handlers available")

func WithIp2GeoHandler(handler Ip2GeoHandler) GeoServiceOption {
	return func(s *GeoService) {
		s.handlers = append(s.handlers, handler)
	}
}

func NewGeoService(loggerImpl common.Logger, opts ...GeoServiceOption) *GeoService {
	service := &GeoService{
		loggerImpl: loggerImpl,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

func (s *GeoService) Ip2Geo(ip string) (models.Ip2GeoHandlerResponse, error) {
	if s.handlers == nil || len(s.handlers) == 0 {
		return models.EmptyIp2GeoHandlerResponse, NoHandlersError
	}

	for i := 0; i < len(s.handlers); i++ {
		handler := s.handlers[i]
		resp, err := handler.Ip2Geo(ip)
		if err == nil {
			s.loggerImpl.Debug("Used handler - %s\n", handler.Name())
			return resp, err
		}
	}

	return models.EmptyIp2GeoHandlerResponse, nil
}

func (s *GeoService) Close() {
	// Cleanup
}
