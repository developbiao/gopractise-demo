package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	// Here's a little snippet that writes . characters to /dev/null
	// using all your cores, and stops afater 10 seconds
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	n := runtime.NumCPU()
	fmt.Printf("Num of CPU is : %d\n", n)
	runtime.GOMAXPROCS(n)

	for i := 0; i < n; i++ {
		go func() {
			for {
				fmt.Fprintf(f, ".")
			}
		}()
	}

	time.Sleep(1 * time.Minute)

	fmt.Println("CPU benchmark is done!")

}
