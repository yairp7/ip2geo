package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/yairp7/ip2geo/internal/common"
	"github.com/yairp7/ip2geo/internal/models"
)

// Ip-Api Handler

type ipApiResponse struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

type IpApiHandler struct {
	loggerImpl common.Logger
}

func NewIpApiHandler(loggerImpl common.Logger) *IpApiHandler {
	return &IpApiHandler{
		loggerImpl: loggerImpl,
	}
}

func (h *IpApiHandler) Name() string {
	return "ip-api.com"
}

func (h *IpApiHandler) Ip2Geo(ip string) (models.Ip2GeoHandlerResponse, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	rawResp, err := http.Get(url)
	if err != nil {
		return models.EmptyIp2GeoHandlerResponse, err
	}

	data, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return models.EmptyIp2GeoHandlerResponse, err
	}
	defer rawResp.Body.Close()

	var response ipApiResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return models.EmptyIp2GeoHandlerResponse, err
	}

	ip2GeoHandlerResponse := models.Ip2GeoHandlerResponse{
		CountryCode: response.CountryCode,
		Lat:         response.Lat,
		Lon:         response.Lon,
	}

	return ip2GeoHandlerResponse, nil
}

// IpInfo Handler

type ipInfoResponse struct {
	Input string             `json:"input"`
	Data  iPInfoResponseData `json:"data"`
}

type iPInfoResponseData struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	LOC      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

type IpInfoHandler struct {
	loggerImpl common.Logger
}

func NewIpInfoHandler(loggerImpl common.Logger) *IpInfoHandler {
	return &IpInfoHandler{
		loggerImpl: loggerImpl,
	}
}

func (h *IpInfoHandler) Name() string {
	return "ipinfo.io"
}

func (h *IpInfoHandler) Ip2Geo(ip string) (models.Ip2GeoHandlerResponse, error) {
	url := fmt.Sprintf("https://ipinfo.io/widget/demo/%s?dataset=geolocation", ip)
	rawResp, err := http.Get(url)
	if err != nil {
		return models.EmptyIp2GeoHandlerResponse, err
	}

	data, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return models.EmptyIp2GeoHandlerResponse, err
	}
	defer rawResp.Body.Close()

	h.loggerImpl.Debug("%v\n", string(data))

	var response ipInfoResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return models.EmptyIp2GeoHandlerResponse, err
	}

	coords := strings.Split(response.Data.LOC, ",")
	lat, _ := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
	lon, _ := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)

	ip2GeoHandlerResponse := models.Ip2GeoHandlerResponse{
		CountryCode: response.Data.Country,
		Lat:         lat,
		Lon:         lon,
	}

	return ip2GeoHandlerResponse, nil
}
