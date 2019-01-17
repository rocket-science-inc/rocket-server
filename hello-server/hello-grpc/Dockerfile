FROM golang:1

RUN apt-get update -yq && apt-get install -yq --no-install-recommends \
    unzip && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
    unzip protoc-3.6.1-linux-x86_64.zip -d protoc3 && \
    mv protoc3/bin/* /usr/local/bin/ && \
    mv protoc3/include/* /usr/local/include/ && \
    ln -s /protoc3/bin/protoc /usr/bin/protoc

RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p ${GOPATH}/src/gitlab.com/pantomath-io/demo-grpc
ADD . ${GOPATH}/src/gitlab.com/pantomath-io/demo-grpc
WORKDIR ${GOPATH}/src/gitlab.com/pantomath-io/demo-grpc

RUN dep ensure

RUN curl -OL https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/third_party/googleapis/google/api/annotations.proto && \
    curl -OL https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/third_party/googleapis/google/api/http.proto && \
    mkdir -p $GOPATH/src/google/api && \
    mv *.proto $GOPATH/src/google/api

RUN protoc -I api/ -I${GOPATH}/src --go_out=plugins=grpc:api api/api.proto

RUN go install ./server && go clean -cache

CMD go run ./server

EXPOSE 7777 7778