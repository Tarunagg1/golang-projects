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

	"github.com/Tarunagg1/student-api/internal/config"
	"github.com/Tarunagg1/student-api/internal/http/handlers/student"
	sqllite "github.com/Tarunagg1/student-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// loggerz

	// db setup

	_, err := sqllite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initialize")

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Server started %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	<-done

	slog.Info("Sutting down the serer")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shut down theserer", slog.String("error", err.Error()))
	}

	slog.Info("Shutting down server successfully")
}
