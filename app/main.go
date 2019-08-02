package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var port = flag.Int("port", 8080, "the port the server will listen on")

func main() {
	flag.Parse()

	r := mux.NewRouter()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", *port),
		Handler:      r,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
