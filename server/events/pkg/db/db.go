package db

import (
	log "github.com/go-kit/kit/log"
)

var logger log.Logger

// Client is an interface defining all methods our custom redis client wrappers must implement
type Client interface {
	// Ping sends a ping to the redis server, this can be used to check the connection
	Ping() error
	// StoreInHset stores value in an hset with the designated key in the designated field
	StoreInHset(key, field string, value []byte) error
	// GetFromHset gets a value fron an hset with the designated key from the designated field
	GetFromHset(key, field string) ([]byte, error)
	// GetMemUsage returns the redis memory usage
	GetMemUsage() (int, error)
	// StartPipe returns a new RedisPipe pipeline
	StartPipe() Pipe
}

type Pipe interface {
	// StoreInHset stores value in an hset with the designated key in the designated field
	StoreInHset(key, field string, value []byte) error
	// GetFromHset gets a value fron an hset with the designated key from the designated field
	GetFromHset(key, field string) ([]byte, error)
	// Executes the pipe and gets the results.
	Execute() ([]byte, error)
}

func NewClient(connectionAddr string) Client {
	return newRedisClient(connectionAddr, "tcp")
}