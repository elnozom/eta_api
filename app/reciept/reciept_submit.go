package reciept

import (
	"encoding/json"
	"eta/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (a *RecieptApi) AuthenticatePOS() (error, *string) {
	var response model.EtaLoginResponse
	// apiUrl := "https://id.preprod.eta.gov.eg"
	resource := "/connect/token"
	data := url.Values{}
	data.Set("client_id", "92fe559b-c17e-4275-a12e-132d34189ef1")
	data.Set("client_secret", "1e0c3a98-b4df-489b-b366-25e3aa5e28c6")
	data.Set("grant_type", "client_credentials")
	u, _ := url.ParseRequestURI(base_url)
	u.Path = resource
	urlStr := u.String()
	fmt.Println(urlStr)
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "https://id.eta.gov.eg/connect/token", strings.NewReader(data.Encode())) // URL-encoded payload
	// r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(r)
	d, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(d, &response)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	return response.AccessToken, nil
	return nil, nil
}
