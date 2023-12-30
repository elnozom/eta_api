package client

import (
	"eta/model"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (c *ApiClient) SetAccessToken(token string) {
	c.accessToken = token
}

type LoginRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (c *ApiClient) Login(req model.EtaAuthunticatePOSRequest) (*model.EtaLoginResponse, error) {
	var resp model.EtaLoginResponse
	var errMsg interface{}
	urlToLogin := fmt.Sprintf("%s/connect/token", c.config.AuthApiUrl)
	response, err := c.client.R().
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     req.ClientId,
			"client_secret": req.ClientSecret,
		}).
		SetHeader("posserial", req.PosSerial).
		SetHeader("pososversion", req.PosOsVersion).
		SetHeader("presharedkey", req.Presharedkey).
		SetErrorResult(&errMsg).
		SetSuccessResult(&resp).
		Post(urlToLogin)
	if err != nil {
		log.Fatal().Err(err).Msg("Login")
	}
	if response.IsErrorState() { // Status code >= 400.
		return nil, err
	}

	log.Debug().Interface("asdrespose", resp).Msg("errrrorrr")
	if response.IsSuccessState() {
		c.SetAccessToken(resp.AccessToken) // Status code is between 200 and 299.
		return &resp, nil
	}

	return &resp, nil
}

func (c *ApiClient) GenerateUUID(req model.Receipt) (*model.GenerateUUIDResponse, error) {
	var resp model.GenerateUUIDResponse
	var errMsg interface{}
	url := fmt.Sprintf("%suuid", c.config.ToolkitApiUrl)
	response, err := c.client.R().
		SetBody(req).
		SetErrorResult(&errMsg).
		SetSuccessResult(&resp).
		Post(url)
	if err != nil {
		log.Fatal().Err(err).Str("url", url).Msg("Login")
	}
	if response.IsErrorState() { // Status code >= 400.
		return nil, err
	}

	log.Debug().Interface("holea", err).Msg("errrrorrr")

	return &resp, nil
}
func (c *ApiClient) SubmitReceitps(loginReq model.EtaAuthunticatePOSRequest, req model.ReceiptSubmitRequest) (*model.InvoiceSubmitResp, error) {
	var resp model.InvoiceSubmitResp
	var errMsg interface{}
	url := fmt.Sprintf("%s/api/v1/receiptsubmissions", c.config.SubmitApiUrl)
	loginResp, err := c.Login(loginReq)
	if err != nil {
		log.Fatal().Err(err).Str("url", url).Msg("Login")
	}
	log.Debug().Interface("loginResp", loginResp).Msg("login response")
	response, err := c.client.R().
		SetBody(req).
		SetBearerAuthToken(loginResp.AccessToken).
		SetErrorResult(&errMsg).
		SetSuccessResult(&resp).
		Post(url)
	if err != nil {
		log.Fatal().Err(err).Str("url", url).Msg("Submit")
	}
	if response.IsErrorState() { // Status code >= 400.
		return nil, err
	}

	return &resp, nil
}
