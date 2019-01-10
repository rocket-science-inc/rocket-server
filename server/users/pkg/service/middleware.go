package service

import (
	"context"
	types "rocket-server/server/users/pkg/types"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UsersService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (users []types.User, error error) {
	defer func() {
		l.logger.Log("method", "Get", "users", users, "error", error)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Add(ctx context.Context, u types.User) (user types.User, error error) {
	defer func() {
		l.logger.Log("method", "Add", "u", u, "user", user, "error", error)
	}()
	return l.next.Add(ctx, u)
}

type authMiddleware struct {
	next UsersService
}

// AuthMiddleware returns a UsersService Middleware.
func AuthMiddleware() Middleware {
	return func(next UsersService) UsersService {
		return &authMiddleware{next}
	}

}
func (a authMiddleware) Get(ctx context.Context) (users []types.User, error error) {
	// Implement your middleware logic here

	return a.next.Get(ctx)
}
func (a authMiddleware) Add(ctx context.Context, u types.User) (user types.User, error error) {
	// Implement your middleware logic here

	return a.next.Add(ctx, u)
}
