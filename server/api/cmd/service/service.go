package service

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"

	log "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
	kit_endpoint "github.com/go-kit/kit/endpoint"
	//opentracing "github.com/go-kit/kit/tracing/opentracing"

	service "rocket-server/server/api/pkg/service"
	endpoint "rocket-server/server/api/pkg/endpoint"

)

var tracer opentracinggo.Tracer
var logger log.Logger

var fs = flag.NewFlagSet("api", flag.ExitOnError)

func Run() {
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("api service", "started")
	defer level.Info(logger).Log("api service", "ended")

	//  Determine which tracer to use.
	level.Info(logger).Log("tracer", "none")
	tracer = opentracinggo.GlobalTracer()

	svc := service.New(getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initCancelInterrupt(g)
	level.Info(logger).Log("exit", g.Run())
}

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	initGRPCHandler(endpoints, g)
	return g
}

func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}

	return
}

func getEndpointMiddleware(logger log.Logger) (mw map[string][]kit_endpoint.Middleware) {
	mw = map[string][]kit_endpoint.Middleware{}

	return
}

func addEndpointMiddlewareToAllMethods(mw map[string][]kit_endpoint.Middleware, m kit_endpoint.Middleware) {
	methods := []string{}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
