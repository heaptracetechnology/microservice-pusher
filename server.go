package main

import (
	"github.com/heaptracetechnology/microservice-pusher/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
