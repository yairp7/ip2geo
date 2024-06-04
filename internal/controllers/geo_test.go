package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/controllers"
	"github.com/yairp7/ip2geo/internal/middlewares"
	"github.com/yairp7/ip2geo/internal/models"
	"github.com/yairp7/ip2geo/internal/services/geo"
	"gopkg.in/stretchr/testify.v1/assert"
)

type testCase struct {
	req        models.Ip2GeoRequest
	res        models.Ip2GeoResponse
	statusCode int
}

var testCases = map[string]testCase{
	"badRequest1": {
		req: models.Ip2GeoRequest{
			ReqID: "1",
		},
		res: models.Ip2GeoResponse{
			ReqID: "1",
			Ip2GeoHandlerResponse: models.Ip2GeoHandlerResponse{
				CountryCode: "UK",
				Lat:         22.4564,
				Lon:         45.2234,
			},
		},
		statusCode: http.StatusBadRequest,
	},
	"badRequest2": {
		req: models.Ip2GeoRequest{
			IP: "123.123.123.123",
		},
		res: models.Ip2GeoResponse{
			ReqID: "1",
			Ip2GeoHandlerResponse: models.Ip2GeoHandlerResponse{
				CountryCode: "UK",
				Lat:         22.4564,
				Lon:         45.2234,
			},
		},
		statusCode: http.StatusBadRequest,
	},
	"goodRequest1": {
		req: models.Ip2GeoRequest{
			ReqID: "1",
			IP:    "123.123.123.123",
		},
		res: models.Ip2GeoResponse{
			ReqID: "1",
			Ip2GeoHandlerResponse: models.Ip2GeoHandlerResponse{
				CountryCode: "UK",
				Lat:         22.4564,
				Lon:         45.2234,
			},
		},
		statusCode: http.StatusOK,
	},
}

var mockLogger = common.NewStdoutLogger(common.DEBUG)

type mockIp2GeoHandler struct{}

func (h *mockIp2GeoHandler) Name() string {
	return "mockIp2GeoHandler"
}

func (h *mockIp2GeoHandler) Ip2Geo(ip string) (models.Ip2GeoHandlerResponse, error) {
	ip2geo := map[string]models.Ip2GeoHandlerResponse{
		"123.123.123.123": {
			CountryCode: "UK",
			Lat:         22.4564,
			Lon:         45.2234,
		},
	}

	if v, ok := ip2geo[ip]; ok {
		return v, nil
	}

	return models.EmptyIp2GeoHandlerResponse, errors.New("bad ip")
}

func setupRouter() *gin.Engine {
	return gin.Default()
}

func Test_GeoController(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()

	geoController := controllers.NewGeoController(
		mockLogger,
		geo.NewGeoService(
			mockLogger,
			geo.WithIp2GeoHandler(&mockIp2GeoHandler{}),
		),
		nil,
	)
	validateIp2GeoReqMiddleware := middlewares.NewValidateIp2GeoRequestMiddleware(mockLogger)
	router.POST("/", validateIp2GeoReqMiddleware.ValidateIp2GeoRequest, geoController.Ip2Geo)

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			mockRequest, _ := json.Marshal(&v.req)
			mockResponse, _ := json.Marshal(&v.res)
			req, _ := http.NewRequest("POST", "/", bytes.NewReader(mockRequest))
			req.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, v.statusCode, w.Code)
			if w.Code == http.StatusOK {
				responseData, _ := io.ReadAll(w.Body)
				assert.Equal(t, string(mockResponse), string(responseData))
			}
		})
	}
}
