package argo

type (
	VersionApi interface {
		GetVersion() (string, error)
	}

	ServerInfo struct {
		Version string
	}
)

func newVersionApi(argo argo) VersionApi {
	return &api{argo}
}

func (api *api) GetVersion() (string, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/version",
		method: "GET",
	})

	if err != nil {
		return "", err
	}

	var result ServerInfo

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return "", err
	}

	return result.Version, nil
}
