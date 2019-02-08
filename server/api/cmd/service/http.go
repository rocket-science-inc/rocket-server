package service

import (
	"net"
	httpgo "net/http"

	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"

	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	httpkit "github.com/go-kit/kit/transport/http"

	endpoint "rocket-server/server/api/pkg/endpoint"
	http "rocket-server/server/api/pkg/transport/http"
)

var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")

func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := getHttpOptions(logger, tracer)

	httpHandler := http.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return httpgo.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

}

func getHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]httpkit.ServerOption {
	options := map[string][]httpkit.ServerOption{
		"AddEvent":  {httpkit.ServerErrorEncoder(http.ErrorEncoder), httpkit.ServerErrorLogger(logger), httpkit.ServerBefore(opentracing.HTTPToContext(tracer, "AddEvent", logger))},
		"GetEvents": {httpkit.ServerErrorEncoder(http.ErrorEncoder), httpkit.ServerErrorLogger(logger), httpkit.ServerBefore(opentracing.HTTPToContext(tracer, "GetEvents", logger))},
	}
	return options
}
