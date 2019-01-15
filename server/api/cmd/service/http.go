package service

import (
	"net"	
	http "net/http"

	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	
	http_kit "github.com/go-kit/kit/transport/http"
	log "github.com/go-kit/kit/log"
	//level "github.com/go-kit/kit/log/level"

	http_handler "rocket-server/server/api/pkg/http/handler"
	endpoint "rocket-server/server/api/pkg/endpoint"
)

var httpAddr = fs.String("http-addr", ":8080", "HTTP listen address")

func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger, tracer)
	// Add your http options here

	httpHandler := http_handler.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

}

func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http_kit.ServerOption {
	options := map[string][]http_kit.ServerOption {

	}
	return options
}