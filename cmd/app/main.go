package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/jjtsksn/cli-calculator/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	done := make(chan struct{})

	go func() {
		defer close(done)
		app.Run(ctx)
	}()

	<-ctx.Done()
	fmt.Printf("\nShutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	select {
	case <-done:
		fmt.Printf("\nApplication stopped gracefully")
	case <-shutdownCtx.Done():
		fmt.Printf("\nShutdown timeout exceeded, forcing exit")
	}
}
