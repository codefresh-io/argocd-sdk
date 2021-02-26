package argo

type (
	ClusterApi interface {
		CreateCluster(clusterOpt ClusterOpt) (*map[string]interface{}, error)
		GetClusters() ([]ClusterItem, error)
	}

	TlsClientConfig struct {
		CaData   string `json:"caData"`
		Insecure bool   `json:"insecure"`
	}

	ClusterConfig struct {
		BearerToken     string          `json:"bearerToken"`
		TlsClientConfig TlsClientConfig `json:"tlsClientConfig"`
	}

	ClusterOpt struct {
		Name   string        `json:"name"`
		Server string        `json:"server"`
		Config ClusterConfig `json:"config"`
	}

	ClusterItem struct {
		Name   string `json:"name"`
		Server string `json:"server"`
	}

	Cluster struct {
		Items []ClusterItem
	}
)

func newClusterApi(argo argo) ClusterApi {
	return &api{argo}
}

func (api *api) GetClusters() ([]ClusterItem, error) {

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/clusters",
		method: "GET",
	})

	if err != nil {
		return nil, err
	}

	var result Cluster

	err = api.argo.decodeResponseInto(resp, &result)

	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

func (api *api) CreateCluster(clusterOpt ClusterOpt) (*map[string]interface{}, error) {

	r := make(map[string]interface{})

	resp, err := api.argo.requestAPI(&requestOptions{
		path:   "/api/v1/clusters",
		method: "POST",
		body:   clusterOpt,
		qs: map[string]string{
			"upsert": "true",
		},
	})

	err = api.argo.decodeResponseInto(resp, &r)

	return &r, err
}
