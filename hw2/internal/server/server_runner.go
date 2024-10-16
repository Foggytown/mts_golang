package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/decode", DecodeHandler)
	mux.HandleFunc("/hard-op", HardOpHandler)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("err trying to listen: %s\n", err)
		}
		fmt.Printf("Stopping server\n")
		ctx.Done()
	}()

	<-ctx.Done()
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("HTTP shutdown error: %v", err)
	}
	fmt.Printf("Graceful shutdown complete.")
}
