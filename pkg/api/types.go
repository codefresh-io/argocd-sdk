package argo

import "net/http"

type (

	// Applications

	ResourceTree struct {
		Nodes []Node
	}

	Node struct {
		Kind   string
		Uid    string
		Health Health
	}

	Health struct {
		Status string `json:"status"`
	}

	ManagedResource struct {
		Items []ManagedResourceItem
	}

	ManagedResourceItem struct {
		Kind        string
		TargetState string
		LiveState   string
		Name        string
	}

	Application struct {
		Items []ApplicationItem
	}

	ArgoApplication struct {
		Status struct {
			Health struct {
				Status string
			}
			Sync struct {
				Status   string
				Revision string
			}
			History        []ApplicationHistoryItem
			OperationState struct {
				FinishedAt string
				SyncResult struct {
					Revision string
				}
			}
		}
		Spec struct {
			Source struct {
				RepoURL string
			}
			Project    string
			SyncPolicy struct {
				Automated interface{}
			}
		}
		Metadata struct {
			Name string
		}
	}

	ApplicationHistoryItem struct {
		Id       int64
		Revision string
	}

	ApplicationItem struct {
		Metadata ApplicationMetadata `json:"metadata"`
		Spec     ApplicationSpec     `json:"spec"`
	}

	ApplicationMetadata struct {
		Name        string `json:"name"`
		UID         string `json:"uid"`
		Namespace   string `json:"namespace"`
		ClusterName string `json:"clusterName"`
	}

	ApplicationSpecDestination struct {
		Server    string `json:"server"`
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	}

	ApplicationSpec struct {
		Project     string                     `json:"project"`
		Destination ApplicationSpecDestination `json:"destination"`
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

	// AuthOptions
	AuthOptions struct {

		// Token - Codefresh token
		Token string
	}

	// Options
	ClientOptions struct {
		Auth   AuthOptions
		Debug  bool
		Host   string
		Client *http.Client
	}

	api struct {
		argo argo
	}

	argo struct {
		token  string
		host   string
		client *http.Client
	}

	requestOptions struct {
		path   string
		method string
		body   interface{}
		qs     map[string]string
	}
)
