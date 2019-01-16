package handler

import (
	"context"
	"encoding/json"
	http "net/http"

	mux "github.com/gorilla/mux"
	handlers "github.com/gorilla/handlers"
	//handler "github.com/99designs/gqlgen/handler"

	http_kit "github.com/go-kit/kit/transport/http"

	endpoint "rocket-server/server/api/pkg/endpoint"
)

// makeGraphqlHandler creates the handler logic
func makeGraphqlHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http_kit.ServerOption) {
	m.Methods("GET","POST","OPTIONS").Path("/graphql").Handler(		
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), 
		)(http_kit.NewServer(endpoints.GraphqlEndpoint, decodeGraphqlRequest, encodeGraphqlResponse, options...)),
	)
}

// decodeGraphqltRequest  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGraphqlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GraphqlRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGraphqlResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGraphqlResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}