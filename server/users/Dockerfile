FROM golang:1

RUN mkdir -p /go/src/rocket-server/server
ADD . /go/src/rocket-server/server
WORKDIR /go/src/rocket-server/server

RUN curl https://glide.sh/get | sh
RUN glide install

RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

CMD watcher -run rocket-server/server/users/cmd -watch rocket-server/server/users
