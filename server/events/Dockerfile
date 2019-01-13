FROM golang:1-alpine

ENV VERSION             "0.0.1-beta"
ENV SERVICE_PATH        rocket-server/server/events

LABEL service.events.version=${VERSION}

RUN apk add --no-cache curl git && \
    curl https://glide.sh/get | sh

RUN mkdir -p /go/src/${SERVICE_PATH}
ADD . /go/src/${SERVICE_PATH}
WORKDIR /go/src/${SERVICE_PATH}

RUN glide install && glide cc
RUN go install ./cmd && go clean -cache

CMD go run ./cmd

EXPOSE 8082