package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				println(err.(string))
			}
		}()
		p()
		wg.Done()
	}()
	wg.Wait()
}

func p() {
	panic("panda foo!")
}
