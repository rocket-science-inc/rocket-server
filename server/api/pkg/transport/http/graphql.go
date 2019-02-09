package http

import (
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	handler "github.com/99designs/gqlgen/handler"

	http "github.com/go-kit/kit/transport/http"

	endpoint "rocket-server/server/api/pkg/endpoint"
	graphql "rocket-server/server/api/pkg/transport/http/graphql"
)

// makePlaygroundHandler creates the handler logic
func makePlaygroundHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}), 
			handlers.AllowedOrigins([]string{"*"}),
		)(handler.Playground("GraphQL playground", "/graphql")),
	)
}

// makeGraphqlHandler creates the handler logic
func makeGraphQLHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET","POST","OPTIONS").Path("/graphql").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET","POST","OPTIONS"}), 
			handlers.AllowedOrigins([]string{"*"}),
		)(handler.GraphQL(graphql.NewExecutableSchema(
				graphql.Config {
					Resolvers: &graphql.Resolver { 
						Endpoints: endpoints,
					},
				},
			),
		)),
	)
}
