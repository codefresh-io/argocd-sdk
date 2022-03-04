package argo

import (
	"errors"
	"fmt"
)

type (
	AuthApi interface {
		UpdatePassword(UpdatePasswordOpt) error
		CheckToken() error
	}

	UpdatePasswordOpt struct {
		CurrentPassword string `json:"currentPassword"`
		UserName        string `json:"name"`
		NewPassword     string `json:"newPassword"`
	}
)

func newAuthApi(argo argo) AuthApi {
	return &api{argo}
}

func (api *api) UpdatePassword(requestOpt UpdatePasswordOpt) error {
	r := make(map[string]interface{})
	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/account/password",
		method: "PUT",
		body:   requestOpt,
	})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = api.argo.decodeResponseInto(resp, &r)
	return err
}

func (api *api) CheckToken() error {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/account",
		method: "GET",
	})

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return err
	}

	if result["error"] != nil {
		return errors.New(fmt.Sprintf("Failed to verify argocd token, reason:  %v", result["error"]))
	}

	return nil
}
