package utils

import (
	"bytes"
	"encoding/json"
	"eta/config"
	"eta/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// client := &http.Client{
//   CheckRedirect: redirectPolicyFunc,
// }

var base_url = "https://api.invoicing.eta.gov.eg/api/v1"

func EtaLogin() (string, error) {
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
}

func SignInvoices(invoices *[]model.Invoice) (*string, error) {
	// var doc model.InvoiceSubmitResp
	jsonValue, _ := json.Marshal(invoices)
	url := fmt.Sprintf("%sSigner", config.Config("SIGNER_URL"))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := string(d)
	resp.Body.Close()
	return &res, nil
}

func SubmitInvoice(authToken *string, document *model.InvoiceSubmitRequest) (*model.EtaSubmitInvoiceResponse, error) {
	client := &http.Client{}
	var response model.EtaSubmitInvoiceResponse

	jsonBody, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}
	// apiUrl := "https://api.preprod.invoicing.eta.gov.eg/api/v1"
	resource := "/documentsubmissions"
	// data.Set("client_id", "c70450b9-5b89-48dd-be15-9cf7629f7dd1")
	// data.Set("client_secret", "7825b824-841c-4f1a-81cd-d8eb60745ee6")
	// data.Set("grant_type", "client_credentials")
	u, _ := url.ParseRequestURI(base_url)
	u.Path = resource
	urlStr := u.String()
	fmt.Println(urlStr)
	r, err := http.NewRequest(http.MethodPost, "https://api.invoicing.eta.gov.eg/api/v1/documentsubmissions", bytes.NewBuffer(jsonBody)) // URL-encoded payload
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", "Bearer "+*authToken)
	r.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(r)
	d, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
