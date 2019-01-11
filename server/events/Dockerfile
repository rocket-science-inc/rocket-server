FROM golang:1

ENV VERSION             "0.0.1-beta"
ENV SERVICE_PATH        rocket-server/server/events

LABEL service.events.version=${VERSION}

RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN curl https://glide.sh/get | sh 

RUN mkdir -p /go/src/${SERVICE_PATH}
ADD . /go/src/${SERVICE_PATH}
WORKDIR /go/src/${SERVICE_PATH}

RUN glide install

CMD watcher -run ${SERVICE_PATH}/cmd -watch ${SERVICE_PATH}

EXPOSE 8082