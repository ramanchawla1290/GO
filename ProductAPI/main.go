package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handlers "github.com/ramanchawla1290/GO/ProductAPI/handlers"

	redoc "github.com/go-openapi/runtime/middleware"
	mux "github.com/gorilla/mux"
)

// Sample REST API
// GORILLA Framework
// JSON Validation with GO VALIDATOR
// Documentation using Swagger and Redoc

func main() {
	// Custom Logger
	logger := log.New(os.Stdout, "REST-API | ", log.LstdFlags|log.Lshortfile)

	// Handler Interface - Implemented using Structs
	prodHandler := handlers.NewProductHandler(logger)

	// GORILLA Router
	r := mux.NewRouter()
	// Middleware for Main Router
	r.Use(prodHandler.MiddlewareLogging)

	// Sub-Router to handle GET Commands only
	getRouter := r.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", handlers.Ping)
	getRouter.HandleFunc("/products", prodHandler.GetAllProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", prodHandler.GetProduct)

	// Handling Documentation using SWAGGER + REDOC middleware
	opts := redoc.RedocOpts{SpecURL: "/swagger.yaml"}
	docsHandler := redoc.Redoc(opts, nil)
	getRouter.Handle("/docs", docsHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Sub-Router to handle POST Commands only
	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/products", prodHandler.AddProduct)

	// Sub-Router to handle PUT Commands only
	putRouter := r.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", prodHandler.UpdateProduct)

	// Middleware for POST & PUT operations
	postRouter.Use(prodHandler.MiddlewareProductValidation)
	putRouter.Use(prodHandler.MiddlewareProductValidation)

	// Sub-Router to handle DELETE Commands only
	deleteRouter := r.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", prodHandler.DeleteProduct)

	// Custom Server
	server := &http.Server{
		Addr:         "localhost:9090",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Printf("Listening to '%s'\n", server.Addr)

	// STARTING CUSTOM SERVER in a GOROUTINE
	go func(server *http.Server) {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}(server)

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, os.Kill)

	serverSignal := <-signalChannel

	logger.Printf("\n\nReceived signal [%v], initiating graceful shutdown\n\n", serverSignal)

	// GRACEFUL SHUTDOWN of the SERVER
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
