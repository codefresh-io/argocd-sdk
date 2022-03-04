package argo

import "errors"

type (
	RepositoryApi interface {
		CreateRepository(CreateRepositoryOpt) error
		GetRepositories() ([]RepositoryItem, error)
	}

	Repository struct {
		Items []RepositoryItem
	}

	RepositoryItem struct {
		Insecure bool   `json:"insecure"`
		Name     string `json:"name"`
		Repo     string `json:"repo"`
		Type     string `json:"type"`
		Username string `json:"username"`
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

func (api *api) GetRepositories() ([]RepositoryItem, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/repositories",
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result Repository

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return result.Items, nil
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

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	err = api.argo.decodeResponseInto(resp, &r)
	return err
}
