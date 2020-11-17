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
)

var (
	port    int
	sigChan = make(chan os.Signal, 1)
)

func main() {
	flag.IntVar(&port, "port", 35005, "The port the server will listen on")
	flag.Parse()

	server := http.Server{
		Addr: fmt.Sprintf(":%v", port),
	}

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigChan
		server.Shutdown(context.Background())
	}()

	log.Print(server.ListenAndServe())
}
