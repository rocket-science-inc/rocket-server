package database

import (

)

type Client interface {
	// Ping sends a ping to the redis server, this can be used to check the connection
	Ping() error
	// StoreInHset stores value in an hset with the designated key in the designated field
	StoreInHset(key, field string, value []byte) error
	// GetFromHset gets a value fron an hset with the designated key from the designated field
	GetFromHset(key, field string) ([]byte, error)
}

func NewClient(connectionAddr string) Client {
	return newGoRedisClient(connectionAddr, "tcp")
}
