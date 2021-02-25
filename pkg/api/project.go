package argo

type (
	ProjectApi interface {
		GetProjects(token string, host string) ([]ProjectItem, error)
	}
)

func newProjectApi(argo argo) ProjectApi {
	return &api{argo}
}

func (api *api) GetProjects(token string, host string) ([]ProjectItem, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/projects",
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	var result Project

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return result.Items, nil
}
