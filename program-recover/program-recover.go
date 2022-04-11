package main

import "fmt"

func main() {

	defer println("defer 1")

	level1()
}

func level1() {
	println("defer func 3")
	defer func() {
		if err := recover(); err != nil {
			println("recovering in progress...")
			fmt.Println("err", err)
		}
	}()

	level2()
}

func level2() {
	defer println("defer func 4")
	panic("foo panda")
}
