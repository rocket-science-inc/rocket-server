FROM golang:1

ENV VERSION "0.0.1-beta"
LABEL service.events.version=${VERSION}

ENV SERVICE_PATH rocket-server/server/events

RUN mkdir -p /go/src/${SERVICE_PATH}
ADD . /go/src/${SERVICE_PATH}
WORKDIR /go/src/${SERVICE_PATH}

RUN curl https://glide.sh/get | sh
RUN glide install

RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

CMD watcher -run ${SERVICE_PATH}/cmd -watch ${SERVICE_PATH}

EXPOSE 8082