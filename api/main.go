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

	"github.com/mick-roper/pr-leaderboard/api/db"
	"github.com/mick-roper/pr-leaderboard/api/routes"
	"github.com/mick-roper/pr-leaderboard/api/types"
)

var (
	port      int
	storeType string
	store     types.Store
	sigChan   = make(chan os.Signal, 1)
)

func main() {
	store = getStore()
	mux := http.NewServeMux()

	routes.ConfigureGithubRoutes(mux, store)

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

	log.Print("Server listening at address ", server.Addr)
	log.Print(server.ListenAndServe())
}

func init() {
	flag.IntVar(&port, "port", 35005, "The port the server will listen on")
	flag.StringVar(&storeType, "store", "memory", "The type of store the application will use")
	flag.Parse()
}

func getStore() types.Store {
	switch storeType {
	default:
		{
			return db.NewMemoryStore()
		}
	}
}
