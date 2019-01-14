package service

import (
	"flag"
	"net"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	google_grpc "google.golang.org/grpc"

	log "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
	kit_endpoint "github.com/go-kit/kit/endpoint"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	kit_grpc "github.com/go-kit/kit/transport/grpc"

	endpoint "rocket-server/server/events/pkg/endpoint"
	grpc "rocket-server/server/events/pkg/grpc/handler"
	pb "rocket-server/server/events/pkg/grpc/pb"
	service "rocket-server/server/events/pkg/service"
)

var tracer opentracinggo.Tracer
var logger log.Logger

var fs = flag.NewFlagSet("events", flag.ExitOnError)
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")

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
	level.Info(logger).Log("service", "started")
	defer level.Info(logger).Log("service", "ended")

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
	initGRPCHandler(endpoints, g)
	return g
}

func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		level.Error(logger).Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		level.Info(logger).Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := google_grpc.NewServer()
		pb.RegisterEventsServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

}

func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]kit_grpc.ServerOption {
	options := map[string][]kit_grpc.ServerOption{
		"AddEvent":  {kit_grpc.ServerErrorLogger(logger), kit_grpc.ServerBefore(opentracing.GRPCToContext(tracer, "AddEvent", logger))},
		"GetEvents": {kit_grpc.ServerErrorLogger(logger), kit_grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetEvents", logger))},
	}
	return options
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
	methods := []string{"GetEvents", "AddEvent"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
