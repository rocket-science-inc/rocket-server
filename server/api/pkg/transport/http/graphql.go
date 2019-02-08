package http

import (
	"context"
	"encoding/json"
	httpgo "net/http"
	
	//handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	handler "github.com/99designs/gqlgen/handler"

	http "github.com/go-kit/kit/transport/http"

	endpoint "rocket-server/server/api/pkg/endpoint"
	graphql "rocket-server/server/api/pkg/transport/http/graphql"
)

// makeGraphqlHandler creates the handler logic
func makeGraphQLHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/").Handler(handler.Playground("GraphQL playground", "/graphql"))
	
	m.Methods("GET","POST","OPTIONS").Path("/graphql").Handler(		
		handler.GraphQL(graphql.NewExecutableSchema(graphql.Config {Resolvers: &graphql.Resolver{}}),
	))

	/*m.Methods("POST").Path("/get-events").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}), 
			handlers.AllowedOrigins([]string{"*"}))
		(http.NewServer(endpoints.GetEventsEndpoint, decodeGetEventsRequest, encodeGetEventsResponse, options...)
	))*/
}

// decodeGraphQLRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGraphQLRequest(_ context.Context, r *httpgo.Request) (interface{}, error) {
	req := endpoint.GetEventsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGraphQLResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGraphQLResponse(ctx context.Context, w httpgo.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
