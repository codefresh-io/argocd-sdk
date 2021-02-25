package argo

import (
	"errors"
	"fmt"
)

type (
	ApplicationApi interface {
		CreateApplication(CreateApplicationOpt) error
		GetApplications(token string, host string) ([]ApplicationItem, error)
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

func (api *api) GetApplications(token string, host string) ([]ApplicationItem, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/applications",
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	var result Application

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

func (api *api) GetResourceTree(applicationName string) (*ResourceTree, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/applications/" + applicationName + "/resource-tree",
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	var result *ResourceTree

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (api *api) GetResourceTreeAll(applicationName string) (interface{}, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/applications/" + applicationName + "/resource-tree",
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	var result interface{}

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return result.(map[string]interface{})["nodes"], nil

}

func (api *api) GetApplication(application string) (map[string]interface{}, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/applications/" + application,
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		// TODO: add error handling and move it to common place
		return nil, errors.New(fmt.Sprintf("Failed to retrieve application, reason %v", resp.Status))
	}

	var result map[string]interface{}

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (api *api) GetManagedResources(applicationName string) (*ManagedResource, error) {
	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/applications/" + applicationName + "/managed-resources",
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	var result ManagedResource

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
