package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/amitpoitrix/go-graphql/config"
	"github.com/amitpoitrix/go-graphql/internal/graph"
	"github.com/amitpoitrix/go-graphql/internal/repository"
	"github.com/amitpoitrix/go-graphql/internal/service"
	"github.com/amitpoitrix/go-graphql/pkg/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize layers
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	resolver := graph.NewResolver(bookService)

	// Create GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	// Setup routes
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", middleware.Logger(srv))

	// Create server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server starting on http://localhost:%s", cfg.Port)
		log.Printf("GraphQL Playground available at http://localhost:%s/", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}