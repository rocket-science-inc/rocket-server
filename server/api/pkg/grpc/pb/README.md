# Protocol Buffers

[Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)

[Go Generated Code](https://developers.google.com/protocol-buffers/docs/reference/go-generated)

### Istall Protocol Buffer and Protocol Generator

```
go get -u github.com/golang/protobuf/proto
GIT_TAG="v1.2.0" # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
```

### Generate Go Protocol Buffers code for Api service

```
protoc api.proto --go_out=plugins=grpc:.
```
