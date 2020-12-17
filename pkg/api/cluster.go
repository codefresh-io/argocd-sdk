package argo

type (
	ClusterApi interface {
		CreateCluster(clusterOpt ClusterOpt) (*map[string]interface{}, error)
	}

	cluster struct {
		argo argo
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
)

func newClusterApi(argo argo) ClusterApi {
	return &cluster{argo}
}

func (cluster *cluster) CreateCluster(clusterOpt ClusterOpt) (*map[string]interface{}, error) {

	r := make(map[string]interface{})

	resp, err := cluster.argo.requestAPI(&requestOptions{
		path:   "/api/v1/clusters",
		method: "POST",
		body:   clusterOpt,
		qs: map[string]string{
			"upsert": "true",
		},
	})

	err = cluster.argo.decodeResponseInto(resp, &r)

	return &r, err
}
