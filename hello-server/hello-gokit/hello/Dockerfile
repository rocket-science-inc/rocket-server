FROM golang

RUN mkdir -p /go/src/rocket-server/hello-server/hello-gokit

ADD . /go/src/rocket-server/hello-server/hello-gokit

RUN curl https://glide.sh/get | sh
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN cd /go/src/rocket-server/hello-server/hello-gokit && glide install

ENTRYPOINT  watcher -run rocket-server/hello-server/hello-gokit/hello/cmd -watch rocket-server/hello-server/hello-gokit/hello
