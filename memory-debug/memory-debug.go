package main

import (
	"log"
	"runtime"
	"time"
)

func testMemory() {
	container := make([]int, 8)

	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println(" ===> loop end.")
}

func main() {
	log.Println("Test Start.")

	testMemory()

	log.Println("force gc.")
	runtime.GC() // Force gc

	log.Println("Done.")

	time.Sleep(3600 * time.Second) //睡眠，保持程序不退出
}
