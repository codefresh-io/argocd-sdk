package argo

import "errors"

type (
	ApplicationApi interface {
		CreateApplication(CreateApplicationOpt) error
	}

	CreateApplicationOpt struct {
		ApiVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`

		Metadata struct {
			Name string `json:"name"`
		} `json:"metadata"`

		Spec struct {
			Project string `json:"project"`

			Destination struct {
				Namespace string `json:"namespace"`
				Server    string `json:"server"`
				Name      string `json:"name"`
			} `json:"destination"`

			Source struct {
				Path           string `json:"path"`
				RepoURL        string `json:"repoURL"`
				TargetRevision string `json:"targetRevision"`
			} `json:"source"`
		} `json:"spec"`
	}
)

func newApplicationApi(argo argo) ApplicationApi {
	return &api{argo}
}

func (api *api) CreateApplication(requestOpt CreateApplicationOpt) error {
	requestOpt.ApiVersion = "argoproj.io/v1alpha1"
	requestOpt.Kind = "Application"

	r := make(map[string]interface{})
	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/applications",
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
