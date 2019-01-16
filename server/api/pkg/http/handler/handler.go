package handler

import (
	"context"
	"encoding/json"
	"errors"
	http "net/http"
	
	//handlers "github.com/gorilla/handlers"
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

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
