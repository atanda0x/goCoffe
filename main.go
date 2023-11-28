package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/atanda0x/goCoffe/handlers"
)

func main() {
	l := log.New(os.Stdout, "Coffee-api", log.LstdFlags)

	// Coffee handler
	ch := handlers.NewCoffee(l)

	sm := http.NewServeMux()
	sm.Handle("/", ch)

	s := &http.Server{
		Addr:         ":8080",           // Configure the bind address
		Handler:      sm,                // Set the default handler
		IdleTimeout:  120 * time.Second, // Max time for connections using TCP keep-Alive
		ReadTimeout:  1 * time.Second,   // Max time to read request from the client
		WriteTimeout: 1 * time.Second,   // Max time to write response to the clent
	}

	// Start the server
	go func() {
		l.Println("Starting server on port :8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error string server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminated, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	s.Shutdown(tc)

}
