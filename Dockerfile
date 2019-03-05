FROM golang

RUN go get github.com/pusher/pusher-http-go

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/microservice-pusher

ADD . /go/src/github.com/heaptracetechnology/microservice-pusher

RUN go install github.com/heaptracetechnology/microservice-pusher

ENTRYPOINT microservice-pusher

EXPOSE 3000