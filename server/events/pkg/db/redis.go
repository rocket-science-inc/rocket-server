package db

import (
	"strconv"
	"strings"

	redis "github.com/mediocregopher/radix.v2/redis"
	log "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
)

// RedigoClient is a wrapper for a "redigo" client. It is not thread safe
type RadixClient struct {
	client  *redis.Client
	network string
}

type RadixPipe struct {
	client *redis.Client
}

// newRedisClient creates a new radix client
func newRedisClient(connectionAddr string, network string, logger log.Logger) Client {
	client, err := redis.Dial(network, connectionAddr)
	if err != nil {
		level.Error(logger).Log("redis", "Failed to create radix client", "error", err)
	}
	return RadixClient {
		client:  client,
		network: network,
	}
}

func (rc RadixClient) Ping() error {
	return rc.client.Cmd("PING").Err
}

func (rc RadixClient) StoreInHset(key, field string, value []byte) error {
	return rc.client.Cmd("HSET", key, field, value).Err
}

func (rp RadixPipe) StoreInHset(key, field string, value []byte) error {
	rp.client.PipeAppend("HSET", key, field, value)
	return nil
}

func (rc RadixClient) GetFromHset(key, field string) ([]byte, error) {
	resp, err := rc.client.Cmd("HGET", key, field).Bytes()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rp RadixPipe) GetFromHset(key, field string) ([]byte, error) {
	rp.client.PipeAppend("HGET", key, field)
	return nil, nil
}

func (rc RadixClient) GetMemUsage() (int, error) {
	resp, err := rc.client.Cmd("INFO", "memory").Str()
	if err != nil {
		return 0, err
	}
	fields := strings.Fields(resp)
	for i := range fields {
		if strings.HasPrefix(fields[i], "used_memory:") {
			return strconv.Atoi(strings.TrimPrefix(fields[i], "used_memory:"))
		}
	}
	return 0, nil
}

func (rp RadixPipe) Execute() ([]byte, error) {
	// Execute the pipe, check if there are errors
	var err error
	for err == nil {
		err = rp.client.PipeResp().Err
	}
	// ignore pipelineempty errors
	if err == redis.ErrPipelineEmpty {
		err = nil
	}
	return nil, err
}

func (rc RadixClient) StartPipe() Pipe {
	return RadixPipe{
		client: rc.client,
	}
}
