package db

import (
	"os"
	"strconv"
	"testing"
	"time"

	log "github.com/go-kit/kit/log"
)

var logger log.Logger

// store up to 10 statements before executing them
const pipelength = 10

// BenchmarkStoreInHsetRedigoTcp benchmarks the StoreInHset function using a radix.v2
// client with a tcp connection
func BenchmarkStoreInHsetRadixTcp(b *testing.B) {
	logger := log.NewLogfmtLogger(os.Stdout)
	client := newRedisClient("localhost:6379", "tcp", logger)
	benchmarkStoreInHset(client, b)
}

// BenchmarkStoreInHsetRedigoUnix benchmarks the StoreInHset function using a radix.v2
// client with a unix socket
func BenchmarkStoreInHsetRadixUnix(b *testing.B) {
	logger := log.NewLogfmtLogger(os.Stdout)
	client := newRedisClient("/tmp/redis.sock", "unix", logger)
	benchmarkStoreInHset(client, b)
}

// BenchmarkStoreInHsetRedigoTcpWithPipe benchmarks the StoreInHset function using a radix.v2
// client with a tcp connection and using a pipe
func BenchmarkStoreInHsetRadixTcpWithPipe(b *testing.B) {
	logger := log.NewLogfmtLogger(os.Stdout)
	client := newRedisClient("localhost:6379", "tcp", logger)
	benchmarkStoreInHsetWithPipe(client, b)
}

// BenchmarkStoreInHsetRedigoUnixWithPipe benchmarks the StoreInHset function using a radix.v2
// client with a unix socket and using a pipe
func BenchmarkStoreInHsetRadixUnixWithPipe(b *testing.B) {
	logger := log.NewLogfmtLogger(os.Stdout)
	client := newRedisClient("/tmp/redis.sock", "unix", logger)
	benchmarkStoreInHsetWithPipe(client, b)
}

func benchmarkStoreInHset(client Client, b *testing.B) {
	// generate 200 bytes of data
	data := make([]byte, 200)
	// key is the current unix timestamp
	key := strconv.FormatInt(time.Now().Unix(), 10)
	// declare the error variable so we only assign inside the benchmark loop, not declare and assign
	var err error
	// maintain a spereate field counter so we dont have field collisions in case the
	// runtime decides to re run the test with the same key somehow
	var fieldCounter int
	for n := 0; n < b.N; n++ {
		err = client.StoreInHset(key, strconv.Itoa(fieldCounter), data)
		if err != nil {
			logger.Log("Error while storing data in HSet", "key", key, "field", fieldCounter, "error", err)
			return
		}
		fieldCounter++
	}
}

func benchmarkStoreInHsetWithPipe(client Client, b *testing.B) {
	// generate 200 bytes of data
	data := make([]byte, 200)
	// key is the current unix timestamp
	key := strconv.FormatInt(time.Now().Unix(), 10)
	// declare the error variable so we only assign inside the benchmark loop, not declare and assign
	var err error
	// maintain a spereate field counter so we dont have field collisions in case the
	// runtime decides to re run the test with the same key somehow
	var fieldCounter, pipeCounter int
	pipe := client.StartPipe()
	for n := 0; n < b.N; n++ {
		err = pipe.StoreInHset(key, strconv.Itoa(fieldCounter), data)
		if err != nil {
			logger.Log("Error while storing data in HSet", "key", key, "field", fieldCounter, "error", err)
			return
		}
		fieldCounter++
		pipeCounter++
		if pipeCounter%pipelength == 0 && n != 0 {
			_, err := pipe.Execute()
			if err != nil {
				logger.Log("Error while executing pipe", "error", err)
				return
			}
			// reset the counter now that we executed the pipe
			pipeCounter = 0
		}
	}
	// If the pipe is not empty, execute the remaining statements.
	if pipeCounter%pipelength == 0 && pipelength == 0 {
		_, err := pipe.Execute()
		if err != nil {
			logger.Log("Error while executing remainder in pipe", "error", err)
			return
		}
	}
}