package service

import (
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(EventsService) EventsService

type loggingMiddleware struct {
	logger log.Logger
	next   EventsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a EventsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next EventsService) EventsService {
		return &loggingMiddleware{logger, next}
	}

}
