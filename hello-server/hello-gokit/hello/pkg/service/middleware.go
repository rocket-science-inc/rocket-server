package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(HelloService) HelloService

type loggingMiddleware struct {
	logger log.Logger
	next   HelloService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a HelloService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next HelloService) HelloService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SayHello(ctx context.Context, s string) (str string, err error) {
	defer func() {
		l.logger.Log("method", "SayHello", "s", s, "str", str, "err", err)
	}()
	return l.next.SayHello(ctx, s)
}

type authMiddleware struct {
	next HelloService
}

// AuthMiddleware returns a HelloService Middleware.
func AuthMiddleware() Middleware {
	return func(next HelloService) HelloService {
		return &authMiddleware{next}
	}

}
func (a authMiddleware) SayHello(ctx context.Context, s string) (str string, err error) {
	// Implement your middleware logic here

	return a.next.SayHello(ctx, s)
}
