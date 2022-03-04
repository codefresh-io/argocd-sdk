package argo

type (
	ProjectApi interface {
		GetProjects() ([]ProjectItem, error)
	}

	Project struct {
		Items []ProjectItem
	}

	ProjectItem struct {
		Metadata ProjectMetadata `json:"metadata"`
	}

	ProjectMetadata struct {
		Name string `json:"name"`
		UID  string `json:"uid"`
	}
)

func newProjectApi(argo argo) ProjectApi {
	return &api{argo}
}

func (api *api) GetProjects() ([]ProjectItem, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/projects",
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result Project

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return result.Items, nil
}
