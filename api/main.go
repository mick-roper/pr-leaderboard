package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mick-roper/pr-leaderboard/api/routes"
)

var (
	port    int
	sigChan = make(chan os.Signal, 1)
)

func main() {
	flag.IntVar(&port, "port", 35005, "The port the server will listen on")
	flag.Parse()

	mux := http.NewServeMux()

	routes.ConfigureGithubRoutes(mux)

	server := http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigChan
		log.Print("Server shutting down...")
		server.Shutdown(context.Background())
	}()

	fmt.Print("Server listening at address", server.Addr)
	log.Print(server.ListenAndServe())
}
