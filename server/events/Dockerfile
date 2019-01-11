FROM golang:1

LABEL service.events.version="0.0.1-beta"

ENV PATH rocket-server/server/events

RUN mkdir -p /go/src/$PATH
ADD . /go/src/$PATH
WORKDIR /go/src/$PATH

RUN curl https://glide.sh/get | sh
RUN glide install

RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

CMD watcher -run $PATH/cmd -watch $PATH