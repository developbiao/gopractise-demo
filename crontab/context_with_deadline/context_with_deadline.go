package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond
func main() {

	// Create deadline
	d := time.Now().Add(shortDuration)

	// Create have deadline context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	// Use select listen 1s which first over
	select {
		case <-time.After(1 * time.Second):
		fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
	}
}
