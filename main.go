package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/atanda0x/goCoffe/docs"
	"github.com/atanda0x/goCoffe/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title classification of Coffee Product API
//
// @Documentation for Coffee Product API222
//
// @SCHEMA: http
// @BasePath: /Coffee/v2
// @Version: 1
//
// @contact.name Atanda Nafiu
// @contact.url https://github.com/atanda0x
// @contact.email atanadakolapo@gmail.com
//
//
// @swagger:meta

func main() {
	l := log.New(os.Stdout, "Coffee-api", log.LstdFlags)

	// Coffee handler
	ch := handlers.NewCoffee(l)

	sm := mux.NewRouter()
	sm.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/Coffee/get", ch.GetCoffees)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/Coffee/update/{id:[0-9]+}", ch.UpdateCoffee)
	putRouter.Use(ch.MiddlewareCoffeeValid)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/Coffee/create", ch.AddCoffe)
	postRouter.Use(ch.MiddlewareCoffeeValid)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/Coffee/delete/{id:[0-9]+}", ch.DeleteCoffee)

	ops := middleware.RedocOpts{SpecURL: "/doc/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// @CORS
	rh := gohandlers.CORS(gohandlers.AllowedOrigin([]string{"http://localhost:8080"}))

	// sm.Handle("/", ch)

	s := &http.Server{
		Addr:         ":8080",           // Configure the bind address
		Handler:      rh(sm),            // Set the default handler
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
