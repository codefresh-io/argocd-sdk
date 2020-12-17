package argo

import "net/http"

type (
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
