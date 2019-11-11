FROM golang

RUN go get github.com/pusher/pusher-http-go

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/oms-services/pusher

ADD . /go/src/github.com/oms-services/pusher

RUN go install github.com/oms-services/pusher

ENTRYPOINT pusher

EXPOSE 3000