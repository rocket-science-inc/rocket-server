package service

import (
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(ApiService) ApiService

type loggingMiddleware struct {
	logger log.Logger
	next   ApiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ApiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ApiService) ApiService {
		return &loggingMiddleware{logger, next}
	}
}
