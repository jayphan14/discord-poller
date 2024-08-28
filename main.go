package main

import (
	"context"
	"discord-poller/poller"
	"discord-poller/util"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5"
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
	Poller       *poller.Poller
	DBConnection *pgx.Conn
}

func New() (*App, error) {
	DBConnectionString, err := util.GetEnv("DATABASE_URL", "", true)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(context.Background(), DBConnectionString)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected To Database")

	// Create poller
	poller, err := poller.New(conn)
	if err != nil {
		return nil, err
	}

	return &App{
		Poller:       poller,
		DBConnection: conn,
	}, nil
}
