package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type key int

const (
	requestID key = iota
	jwt
)

func status(w http.ResponseWriter, req *http.Request) {
	// Add reqeust id to context
	ctx := context.WithValue(req.Context(), requestID, uuid.NewV4().String())
	// Add credentials to content
	ctx = context.WithValue(ctx, jwt, req.Header.Get("Authorization"))

	upDB, err := isDatabaseUp(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	upAuth, err := isMonitoringUp(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "DB up: %t | Monitoring up: %t\n", upDB, upAuth)
}

func isDatabaseUp(ctx context.Context) (bool, error) {
	// Retrieve the reqeust ID value
	reqID, ok := ctx.Value(requestID).(string)
	if !ok {
		return false, fmt.Errorf("RequestID in context does not have the expected type")
	}
	log.Printf("req %s- checking db status", reqID)
	return true, nil
}

func isMonitoringUp(ctx context.Context) (bool, error) {
	// Retrieve the request ID value
	reqID, ok := ctx.Value(requestID).(string)
	if !ok {
		return false, fmt.Errorf("RequestID in context does not have the expected type")
	}
	log.Printf("req %s- checking monitoring status", reqID)
	return true, nil
}

func main() {
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Main func done!")
}
