package service

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	prometheus "github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"

	endpointkit "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
	prometheuskit "github.com/go-kit/kit/metrics/prometheus"

	endpoint "rocket-server/server/api/pkg/endpoint"
	service "rocket-server/server/api/pkg/service"
)

var tracer opentracinggo.Tracer
var logger log.Logger

var fs = flag.NewFlagSet("api", flag.ExitOnError)
var debugAddr = fs.String("debug.addr", ":8080", "Debug and metrics listen address")
var jaegerURL = fs.String("jaeger-url", "", "Enable Jaeger tracing via a collector URL")

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

	// Determine which tracer to use.
	if *jaegerURL != "" {
		// TODO: add Jaeger tracer
		logger.Log("tracer", "Jaeger", "URL", *jaegerURL)
	} else {
		logger.Log("tracer", "none")
		tracer = opentracinggo.GlobalTracer()
	}

	// Create service
	svc := service.New(getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	group := createService(eps)
	// Register metrics
	initMetricsEndpoint(group)

	// Register shutdown
	initCancelInterrupt(group)
	level.Info(logger).Log("exit", group.Run())
}

func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(logger, mw)
	return mw
}

func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}

func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpointkit.Middleware) {
	mw = map[string][]endpointkit.Middleware{}
	duration := prometheuskit.NewSummaryFrom(prometheus.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "events",
	}, []string{"method", "success"})
	addDefaultEndpointMiddleware(logger, duration, mw)
	return mw
}

func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheuskit.Summary, mw map[string][]endpointkit.Middleware) {
	mw["GetEvents"] = []endpointkit.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetEvents")), endpoint.InstrumentingMiddleware(duration.With("method", "GetEvents"))}
	mw["AddEvent"] = []endpointkit.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "AddEvent")), endpoint.InstrumentingMiddleware(duration.With("method", "AddEvent"))}
}

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}

func initMetricsEndpoint(g *group.Group) {
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http.Serve(debugListener, http.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
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
