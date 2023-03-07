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
)

func sayHi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Girl!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/say", sayHi)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Initialization server goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// Wait fro interrupt singal to gracefully shutdown the server with
	// a time of 5 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Wait Shutdown server...")

	// The context is used to information the server it has 3 seconds to fisihe
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//  Simulation do someting timeout
	time.Sleep(5 * time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
