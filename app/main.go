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
var githubKey = flag.String("github-key", "", "the key that should be used to query the github APIs")
var repos = flag.String("github-repos", "", "the repos that should be interrogated")

func main() {
	flag.Parse()

	if *githubKey == "" {
		log.Fatal("a github key must be provided")
	}

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", *port),
		Handler:      r,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Println("server listening on port", *port)
	log.Fatal(server.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("404 file not found"))
}
