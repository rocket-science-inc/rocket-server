FROM golang:1

# setup environment
ENV VERSION             "0.0.2-beta"
ENV SERVICE_PATH        rocket-server/server/api

LABEL service.api.version=${VERSION}

# install system dependecies
RUN curl https://glide.sh/get | sh

# create workspace and get project sources
RUN mkdir -p /go/src/${SERVICE_PATH}
ADD . /go/src/${SERVICE_PATH}
WORKDIR /go/src/${SERVICE_PATH}

# install project dependecies
RUN glide install && glide cc
# compile
RUN go install ./cmd && go clean -cache
# test
RUN go test -v ./... && go clean -testcache
# run
CMD go run ./cmd

# network
EXPOSE 8080 8081