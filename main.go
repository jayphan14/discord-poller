package main

import (
	"discord-poller/poller"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app, err := New()
	if err != nil {
		fmt.Println("Error starting app:", err)
		return
	}

	go app.Poller.Poll()

	// Graceful shutdown handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Shutting down...")
}

type App struct {
	DataStore map[string]bool
	Poller    *poller.Poller
}

func New() (*App, error) {
	newDataStore := map[string]bool{}
	poller, err := poller.New(&newDataStore)

	if err != nil {
		return nil, err
	}

	return &App{
		DataStore: newDataStore,
		Poller:    poller,
	}, nil
}
