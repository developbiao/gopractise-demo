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

func main() {
	// Create the notification channel
	bye := make(chan os.Signal)
	signal.Notify(bye, os.Interrupt, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.Handle("/status", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "OK")
		},
	))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Lanuch the server in another goroutine
	go func() {
		// Lanuch the server
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server :%q\n", err)
		}
	}()

	// Wait for os singal
	sig := <-bye
	// The code below is executed when we recevied an os.Singal
	log.Printf("Detected os signal %s \n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	err := srv.Shutdown(ctx)
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Main func done!")
}
