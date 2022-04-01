package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

// Execute a command with goroutine executing 2 seconds
// 1 seconds will be kill command
func main() {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	// Create result queue
	resultChan = make(chan *result, 1000)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)

		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo hello;")

		// Executing task capture output
		output, err = cmd.CombinedOutput()

		// Give executed result for main
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)

	cancelFunc()

	// Main waiting child exit and print result
	res = <-resultChan

	// Print result
	fmt.Println(res.err, string(res.output))

}
