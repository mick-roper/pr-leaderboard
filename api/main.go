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

	"github.com/mick-roper/pr-leaderboard/api/auth"
	"github.com/mick-roper/pr-leaderboard/api/db"
	"github.com/mick-roper/pr-leaderboard/api/routes"
	"github.com/mick-roper/pr-leaderboard/api/types"
)

var (
	port        int
	storeType   string
	dataStore   types.Store
	apiKeyStore auth.APIKeyStore

	sigChan = make(chan os.Signal, 1)
)

func main() {
	mux := http.NewServeMux()

	routes.ConfigureGithubRoutes(mux, dataStore)
	routes.ConfigureAPIRoutes(mux, dataStore, apiKeyStore)

	server := http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigChan
		log.Println("Server shutting down...")
		server.Shutdown(context.Background())
	}()

	log.Println("Server listening at address ", server.Addr)
	log.Println(server.ListenAndServe())
}

func init() {
	flag.IntVar(&port, "port", 35005, "The port the server will listen on")
	flag.StringVar(&storeType, "store", "memory", "The type of store the application will use")
	flag.Parse()

	log.Println("Using store type: ", storeType)

	dataStore = getDataStore()
	apiKeyStore = getAPIKeyStore()
}

func getDataStore() types.Store {
	switch storeType {
	case "redis":
		{
			redisAddress := os.Getenv("REDIS_ADDRESS")

			if redisAddress == "" {
				log.Fatal("REDIS_ADDRESS env var has not been set")
			}

			s, err := db.NewRedisStore(redisAddress)

			if err != nil {
				log.Fatal(err)
			}

			return s
		}
	default:
		{
			return db.NewMemoryStore()
		}
	}
}

func getAPIKeyStore() auth.APIKeyStore {
	store, err := auth.NewSimpleAPIKeyStore("abc-123")
	if err != nil {
		log.Fatal(err)
	}

	return store
}
