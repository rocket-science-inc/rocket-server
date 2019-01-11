package db

import (
	"fmt"
)

// Client is an interface defining all methods our custom redis client wrappers must implement
type Client interface {
	// stringer interface is used to format the client when we print
	fmt.Stringer
	// Ping sends a ping to the redis server, this can be used to check the connection
	Ping() error
	// Store stores value with the designated key in the designated field
	Store(key, field string, value []byte) error
	// GetFromHset gets a value with the designated key from the designated field
	Get(key, field string) ([]byte, error)
	// GetMemUsage returns the client memory usage
	GetMemUsage() (int, error)
	// StartPipe returns a new Pipe pipeline
	//StartPipe() Pipe
}

type ConnectionType int
const (
	Tcp ConnectionType = iota
	Unix
)

func NewClient(connectionAddr string, conType ConnectionType) Client {
	return newRedisClient(connectionAddr, conType)
}

// connTypeToString returns a string representation of a ConnectionType
func connTypeToString(conType ConnectionType) string {
	var network string
	switch conType {
	case Tcp:
		network = "tcp"
		break
	case Unix:
		network = "unix"
		break
	default:
		network = "tcp"
	}
	return network
}