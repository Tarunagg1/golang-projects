package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"githug.com/tarunagg1/student-api/internal/config"
	student "githug.com/tarunagg1/student-api/internal/http/handlers"
)

func main() {
	fmt.Println("here")
	// load config
	cfg := config.MustLoad()

	// databse setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server
	server := http.Server{
		Addr:         cfg.Addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second, // Read timeout for incoming requests
		WriteTimeout: 10 * time.Second, // Write timeout for server responses
		IdleTimeout:  15 * time.Second, // Idle timeout for keep-alive connections
	}

	slog.Info("Server started %s", slog.String("address", cfg.Addr))
	fmt.Printf("Server started %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to status sever")
		}
	}()

	<-done

	slog.Info("Sutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shut down the server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}

// go run .\cmd\student-api\main.go -config config/local.yaml
