package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")
	if err != nil {
		var pErr *os.PathError
		if errors.As(err, &pErr) {
			fmt.Println("Faile to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")

}
