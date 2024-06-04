package models

type Ip2GeoHandlerResponse struct {
	CountryCode string  `json:"countryCode"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

type Ip2GeoResponse struct {
	ReqID string `json:"reqId"`
	Ip2GeoHandlerResponse
}

var EmptyIp2GeoHandlerResponse Ip2GeoHandlerResponse = Ip2GeoHandlerResponse{}
