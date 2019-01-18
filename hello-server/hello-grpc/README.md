# Hello ~~World~~ gRPC

[How we use gRPC to build a client/server system in Go](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2)


## Build

```
docker build --rm -t hello-grpc .
```

## Run Server

```
docker run -d --rm --name grpc-server -p 7777:7777 -p 8080:7778 hello-grpc 
```

## Kill Server

```
docker kill grpc-server
```

## Run Client

```
docker run -it --rm --name grpc-client hello-grpc go run ./client
```
