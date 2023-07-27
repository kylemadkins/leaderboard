package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kylemadkins/leaderboard/db"
	"github.com/kylemadkins/leaderboard/handlers"
)

func main() {
	addr := "8000"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	database, err := db.Initialize(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to set up database: %s", err.Error())
	}
	defer database.Conn.Close()

	httpHandler := handlers.NewHandler(database)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Listening on port %s...", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server...")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Unable to shut server down correctly: %s\n", err.Error())
		os.Exit(1)
	}
}
