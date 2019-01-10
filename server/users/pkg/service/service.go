package service

import (
	log "github.com/Sirupsen/logrus"

	"context"
	"rocket-server/server/users/pkg/types"
	"rocket-server/server/users/pkg/db"
)

type UsersService interface {
	Get(ctx context.Context) (users []types.User, error error)
	Add(ctx context.Context, u types.User) (user types.User, error error)
}

type basicUsersService struct{}

func (b *basicUsersService) Get(ctx context.Context) (users []types.User, error error) {
	dbConnectionString := "localhost:6379"
	log.Debug("Connect to redis server at address ", dbConnectionString)

	db.NewClient(db.Redis, dbConnectionString, db.Unix)
	// TODO implement the business logic of Get

	return users, error
}
func (b *basicUsersService) Add(ctx context.Context, u types.User) (user types.User, error error) {
	// TODO implement the business logic of Add
	return user, error
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	return &basicUsersService{}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
