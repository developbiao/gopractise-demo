package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	rootCtx := context.Background()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
	// Create context
	ctx, cancel := context.WithTimeout(rootCtx, 50*time.Millisecond)
	defer cancel()
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println("Resp recievied", resp)
	fmt.Println("Main func done!")
}
