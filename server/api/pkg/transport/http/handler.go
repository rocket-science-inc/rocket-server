package http

import (
	"context"
	"encoding/json"
	"errors"
	httpgo "net/http"

	mux "github.com/gorilla/mux"

	http "github.com/go-kit/kit/transport/http"

	endpoint "rocket-server/server/api/pkg/endpoint"
)

//  NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) httpgo.Handler {
	m := mux.NewRouter()
	makeGraphQLHandler(m, endpoints, options["GraphQL"])
	makePlaygroundHandler(m, endpoints, options["Playground"])
	return m
}

func ErrorEncoder(_ context.Context, err error, w httpgo.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func ErrorDecoder(r *httpgo.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return httpgo.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
