package model

type ListOrdersRequest struct {
	Store       int `query:"store"`
	Status      int `query:"status"`
	TransSerial int `query:"transSerial"`
}
