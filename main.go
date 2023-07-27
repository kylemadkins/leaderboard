package main

import (
	"context"
	"fmt"
	"log"
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

	database, err := db.Initialize(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to set up database: %s", err.Error())
	}
	defer database.Conn.Close()

	httpHandler := handlers.NewHandler(database)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", addr),
		Handler: httpHandler,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(err.Error())
		}
	}()
	log.Printf("Server started on port %s\n", addr)
	<-done
	log.Println("Exiting server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to exit: %+v", err)
	}
}
