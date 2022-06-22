package utils

import (
	"bytes"
	"encoding/json"
	"eta/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// client := &http.Client{
// 	CheckRedirect: redirectPolicyFunc,
// }

var base_url = "https://api.preprod.invoicing.eta.gov.eg/api/v1/"

func EtaLogin() (string, error) {
	var response model.EtaLoginResponse
	// apiUrl := "https://id.preprod.eta.gov.eg"
	resource := "/connect/token"
	data := url.Values{}
	data.Set("client_id", "c70450b9-5b89-48dd-be15-9cf7629f7dd1")
	data.Set("client_secret", "7825b824-841c-4f1a-81cd-d8eb60745ee6")
	data.Set("grant_type", "client_credentials")
	u, _ := url.ParseRequestURI(base_url)
	u.Path = resource
	urlStr := u.String()
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(r)
	d, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(d, &response)
	resp.Body.Close()
	return response.AccessToken, nil
}

func SignDocument(document *model.Invoice) error {
	client := &http.Client{}
	apiUrl := "http://localhost:5000/"
	resource := "/sign"
	data := []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	fmt.Println(data)
	// data.Set("client_id", "c70450b9-5b89-48dd-be15-9cf7629f7dd1")
	// data.Set("client_secret", "7825b824-841c-4f1a-81cd-d8eb60745ee6")
	// data.Set("grant_type", "client_credentials")
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()
	r, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(data)) // URL-encoded payload
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(r)
	d, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(d, &document)
	resp.Body.Close()
	return nil
}

func SubmitInvoice(authToken *string) (string, error) {
	client := &http.Client{}
	var response model.EtaSubmitInvoiceResponse
	// apiUrl := "https://api.preprod.invoicing.eta.gov.eg/api/v1"
	resource := "/documentsubmissions"
	data := []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	fmt.Println(data)
	// data.Set("client_id", "c70450b9-5b89-48dd-be15-9cf7629f7dd1")
	// data.Set("client_secret", "7825b824-841c-4f1a-81cd-d8eb60745ee6")
	// data.Set("grant_type", "client_credentials")
	u, _ := url.ParseRequestURI(base_url)
	u.Path = resource
	urlStr := u.String()
	r, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(data)) // URL-encoded payload
	if err != nil {
		return "", err
	}
	r.Header.Add("Authorization", "Bearer "+*authToken)
	r.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(r)
	d, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(d, &response)
	resp.Body.Close()
	return response.SubmissionId, nil
}
