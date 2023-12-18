package reciept

import (
	"eta/config"
)

type RecieptApiInterface interface {
	AuthenticatePOS() (error, *string)
}

type RecieptApi struct {
	config config.Config
}

func NewRecieptApi(config config.Config) RecieptApiInterface {
	return &RecieptApi{
		config: config,
	}
}
