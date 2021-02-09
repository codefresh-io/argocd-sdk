package argo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type (
	Argo interface {
		requestAPI(*requestOptions) (*http.Response, error)
		Clusters() ClusterApi
		Auth() AuthApi
		Repository() RepositoryApi
		Application() ApplicationApi
	}
)

func GetToken(username string, password string, host string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}

	message := map[string]interface{}{
		"username": username,
		"password": password,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		return "", errors.New("application error, cant retrieve argo token")
	}

	resp, err := client.Post(host+"/api/v1/session", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 401 {
		return "", errors.New("cant retrieve argocd token, permission denied")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return result["token"].(string), nil
}

func New(opt *ClientOptions) Argo {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	httpClient := &http.Client{}
	if opt.Client != nil {
		httpClient = opt.Client
	}

	return &argo{
		host:   opt.Host,
		token:  opt.Auth.Token,
		client: httpClient,
	}
}

func (a argo) Clusters() ClusterApi {
	return newClusterApi(a)
}

func (a argo) Auth() AuthApi {
	return newAuthApi(a)
}

func (a argo) Repository() RepositoryApi {
	return newRepositoryApi(a)
}

func (a argo) Application() ApplicationApi {
	return newApplicationApi(a)
}

func (a argo) requestAPI(opt *requestOptions) (*http.Response, error) {
	var body []byte
	finalURL := fmt.Sprintf("%s%s", a.host, opt.path)
	if opt.qs != nil {
		finalURL += toQS(opt.qs)
	}
	if opt.body != nil {
		body, _ = json.Marshal(opt.body)
	}
	request, err := http.NewRequest(opt.method, finalURL, bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+a.token)
	request.Header.Set("Content-Type", "application/json")

	response, err := a.client.Do(request)
	if err != nil {
		return response, err
	}
	return response, nil
}

func toQS(qs map[string]string) string {
	var arr = []string{}
	for k, v := range qs {
		arr = append(arr, fmt.Sprintf("%s=%s", k, v))
	}
	return "?" + strings.Join(arr, "&")
}

func (c *argo) decodeResponseInto(resp *http.Response, target interface{}) error {
	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *argo) getBodyAsString(resp *http.Response) (string, error) {
	body, err := c.getBodyAsBytes(resp)
	return string(body), err
}

func (c *argo) getBodyAsBytes(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
