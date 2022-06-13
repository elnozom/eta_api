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

func EtaLogin() (string, error) {
	var response model.EtaLoginResponse
	apiUrl := "https://id.preprod.eta.gov.eg"
	resource := "/connect/token"
	data := url.Values{}
	data.Set("client_id", "c70450b9-5b89-48dd-be15-9cf7629f7dd1")
	data.Set("client_secret", "7825b824-841c-4f1a-81cd-d8eb60745ee6")
	data.Set("grant_type", "client_credentials")
	u, _ := url.ParseRequestURI(apiUrl)
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

func SubmitInvoice() (string, error) {
	client := &http.Client{}
	var response model.EtaSubmitInvoiceResponse
	apiUrl := "https://api.preprod.invoicing.eta.gov.eg/api/v1"
	resource := "/documentsubmissions"
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
		return "", err
	}
	r.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk2RjNBNjU2OEFEQzY0MzZDNjVBNDg1MUQ5REM0NTlFQTlCM0I1NTQiLCJ0eXAiOiJhdCtqd3QiLCJ4NXQiOiJsdk9tVm9yY1pEYkdXa2hSMmR4Rm5xbXp0VlEifQ.eyJuYmYiOjE2NTUxMjcwOTcsImV4cCI6MTY1NTEzMDY5NywiaXNzIjoiaHR0cHM6Ly9pZC5wcmVwcm9kLmV0YS5nb3YuZWciLCJhdWQiOiJJbnZvaWNpbmdBUEkiLCJjbGllbnRfaWQiOiJjNzA0NTBiOS01Yjg5LTQ4ZGQtYmUxNS05Y2Y3NjI5ZjdkZDEiLCJJc1RheFJlcHJlcyI6IjEiLCJJc0ludGVybWVkaWFyeSI6IjAiLCJJbnRlcm1lZElkIjoiMCIsIkludGVybWVkUklOIjoiIiwiSW50ZXJtZWRFbmZvcmNlZCI6IjIiLCJuYW1lIjoiMjg4MjcxOTk4OmM3MDQ1MGI5LTViODktNDhkZC1iZTE1LTljZjc2MjlmN2RkMSIsIlNTSWQiOiJmZTRmZjU4NS1mNDQ4LWRlYmEtNjQwMC00ZDQzOTNkY2EyNWMiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJub3pvbSIsIlRheElkIjoiMjU0ODE4IiwiVGF4UmluIjoiMjg4MjcxOTk4IiwiUHJvZklkIjoiMjQxNTc2IiwiSXNUYXhBZG1pbiI6IjAiLCJJc1N5c3RlbSI6IjEiLCJOYXRJZCI6IiIsIlRheFByb2ZUYWdzIjpbIkIyQiIsIkIyQyJdLCJzY29wZSI6WyJJbnZvaWNpbmdBUEkiXX0.SxZgfOzpJQNqsJD3kbL6Cbv4PHxApunzIR3WWY8YIP8ZAVbi2MiOmrEpxWEnngBP3HqgIsNlarVzcXjfC6xz1ENSQ8Bu-IXdWur65Dde7ZjynxNC8Kryxho0l-H1YAfCzDCuHKzfuj0oO1_50tYSuogZGqx4AVvHpfLpg3a3b2bQP1oZVajqGZhMh6_AyL_0p54Uq1eMyz_cMveviJVTAlG2oyf8uEXQ-zsps3OsUjBeZFcKD3xndqZzi5aLPFF3K6G7K6Jg8YS3PbAMtqaJvCcfW4YIsJezJyoz9SYYP2eG0444w1JMOwm_HsIiy7M4-1gVW2EhzImzIXqFTHbjcg")
	r.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(r)
	d, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(d, &response)
	resp.Body.Close()
	return response.SubmissionId, nil
}
