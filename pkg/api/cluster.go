package api

import (
	"github.com/fatih/structs"
)

type (
	ClusterApi interface {
		CreateCluster(clusterOpt ClusterOpt) error
	}

	cluster struct {
		argo argo
	}

	ClusterConfig struct {
		BearerToken     string `json:"bearerToken"`
		TlsClientConfig struct {
			CaData   string `json:"caData"`
			Insecure bool   `json:"insecure"`
		} `json:"tlsClientConfig"`
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

func (cluster *cluster) CreateCluster(clusterOpt ClusterOpt) error {

	_, err := cluster.argo.requestAPI(&requestOptions{
		path:   "/api/v1/clusters",
		method: "POST",
		body:   structs.Map(clusterOpt),
		qs: map[string]string{
			"upsert": "true",
		},
	})

	return err
}
