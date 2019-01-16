package handler

import (
	http "net/http"

	mux "github.com/gorilla/mux"

	http_kit "github.com/go-kit/kit/transport/http"

	endpoint "rocket-server/server/api/pkg/endpoint"
)

//  NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http_kit.ServerOption) http.Handler {
	m := mux.NewRouter()
	makeGraphqlHandler(m, endpoints, options["Graphql"])
	return m
}