package argo

import "errors"

type (
	RepositoryApi interface {
		CreateRepository(CreateRepositoryOpt) error
	}

	CreateRepositoryOpt struct {
		Repo          string `json:"repo"`
		Username      string `json:"username"`
		Password      string `json:"password"`
		SshPrivateKey string `json:"sshPrivateKey"`

		Insecure bool   `json:"insecure"`
		Name     string `json:"name"` // for helm repos only
		Type     string `json:"type"` // git or helm
	}
)

func newRepositoryApi(argo argo) RepositoryApi {
	return &api{argo}
}

func (api *api) CreateRepository(requestOpt CreateRepositoryOpt) error {
	r := make(map[string]interface{})
	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/repositories",
		method: "POST",
		body:   requestOpt,
	})
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	err = api.argo.decodeResponseInto(resp, &r)
	return err
}
