package argo

type (
	AuthApi interface {
		UpdatePassword(UpdatePasswordOpt) error
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
	err = api.argo.decodeResponseInto(resp, &r)
	return err
}
