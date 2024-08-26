package main

import (
	"discord-poller/poller"
	"fmt"
)

func main() {
	app, err := New()
	if err != nil {
		fmt.Println("Error starting app:", err)
		return
	}

	go app.Poller.Poll()
	select {}
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
