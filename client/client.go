package client

import (
	"eta/config"
	"eta/model"

	"github.com/imroc/req/v3"
)

type ApiClientInterface interface {
	SetAccessToken(token string)
	Login() (*model.EtaLoginResponse, error)
	// Login() (*types.LoginResponse, error)
	// DoProccessPaymentReuest(req types.ProccessPaymentRequest, token string) (*req.Response, *types.ProccessPaymentResponse, *ErrorResponse)
	// DecorateProcessPaymentRequest(req *types.ProccessPaymentRequest)
	// ProccessPayment(req types.ProccessPaymentRequest) (*types.ProccessPaymentResponse, error)
}

type ApiClient struct {
	client      *req.Client
	config      *config.Config
	accessToken string
}

func NewApiClient(config *config.Config) ApiClientInterface {
	c := req.C().DevMode()
	return &ApiClient{
		client:      c,
		config:      config,
		accessToken: "",
	}
}
