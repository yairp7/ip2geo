package models

type Ip2GeoRequest struct {
	ReqID string `json:"reqId" validate:"required"`
	IP    string `json:"ip" validate:"required"`
}
