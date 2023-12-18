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

func (c *ApiClient) Login() (*model.EtaLoginResponse, error) {
	var resp model.EtaLoginResponse
	var errMsg interface{}
	urlToLogin := fmt.Sprintf("%s/connect/token", c.config.AuthApiUrl)
	response, err := c.client.R().
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     c.config.POSClientID,
			"client_secret": c.config.POSClientSecret,
		}).
		SetHeader("posserial", c.config.PosSerial).
		SetHeader("pososversion", c.config.PosOsVersion).
		SetHeader("presharedkey", "").
		SetErrorResult(&errMsg).
		Post(urlToLogin)
	if err != nil {
		log.Fatal().Err(err).Msg("Login")
	}
	if response.IsErrorState() { // Status code >= 400.
		return nil, err
	}

	log.Debug().Interface("holea", err).Msg("errrrorrr")
	if response.IsSuccessState() {
		c.SetAccessToken(resp.AccessToken) // Status code is between 200 and 299.
		return &resp, nil
	}

	return &resp, nil
}
