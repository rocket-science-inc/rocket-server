package database

import (
	redis "github.com/go-redis/redis"
)

type GoRedisClient struct {
	client  *redis.Client
	network string
}

// Creates a go-redis client
func newGoRedisClient(connectionAddr string, network string) Client {
	client := redis.NewClient(&redis.Options{
		Network:  network,
		Addr:     connectionAddr,
		Password: "",
		DB:       0,
	})
	return GoRedisClient {
		client:  client,
		network: network,
	}
}

func (rc GoRedisClient) Ping() error {
	return rc.client.Ping().Err()
}

func (rc GoRedisClient) StoreInHset(key, field string, value []byte) error {
	return rc.client.HSet(key, field, value).Err()
}

func (rc GoRedisClient) GetFromHset(key, field string) ([]byte, error) {
	cmd := rc.client.HGet(key, field)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return []byte(cmd.Val()), nil
}
