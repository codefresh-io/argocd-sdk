package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type (
	Argo interface {
		requestAPI(*requestOptions) (*http.Response, error)
		Clusters() ClusterApi
	}
)

func New(opt *ClientOptions) Argo {
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
	request.Header.Set("Authorization", a.token)
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
