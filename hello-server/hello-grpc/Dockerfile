FROM golang:1

RUN apt-get update -yq && apt-get install -yq --no-install-recommends \
    unzip && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
    unzip protoc-3.6.1-linux-x86_64.zip -d protoc3 && \
    mv protoc3/bin/* /usr/local/bin/ && \
    mv protoc3/include/* /usr/local/include/ && \
    ln -s /protoc3/bin/protoc /usr/bin/protoc

RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

RUN go get -u github.com/golang/protobuf/proto
RUN git -C ${GOPATH}/src/github.com/golang/protobuf checkout "v1.2.0"
RUN go install github.com/golang/protobuf/protoc-gen-go

WORKDIR ${GOPATH}/src/gitlab.com/pantomath-io/demo-grpc
ADD . .

RUN make

CMD go run ./server

EXPOSE 7777 7778