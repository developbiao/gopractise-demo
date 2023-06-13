package main

import "io"

type myWriter struct {

}

func (w myWriter) Write(p []byte) (n int, err error) {
	return
}

func main() {
	// Check *myWriter type is implement io.Writer interface
	var _ io.Writer = (*myWriter)(nil)

	// Check *myWriter type is implement io.Writer interface
	var _ io.Writer = myWriter{}

}