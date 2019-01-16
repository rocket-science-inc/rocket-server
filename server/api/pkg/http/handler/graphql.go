package handler

import (
	mux "github.com/gorilla/mux"
	handler "github.com/99designs/gqlgen/handler"

	http_kit "github.com/go-kit/kit/transport/http"

	endpoint "rocket-server/server/api/pkg/endpoint"
	"rocket-server/server/api/pkg/http/graphql"
)

// makeGraphqlHandler creates the handler logic
func makeGraphqlHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http_kit.ServerOption) {
	m.Methods("GET").Path("/").Handler(handler.Playground("GraphQL playground", "/graphql"))
	m.Methods("GET","POST","OPTIONS").Path("/graphql").Handler(		
		handler.GraphQL(graphql.NewExecutableSchema(graphql.Config {Resolvers: &graphql.Resolver{}}),
	))
}