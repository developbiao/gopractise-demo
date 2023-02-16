package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	writeLength, err := f.WriteString("你好，中国🫡\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(writeLength, "Bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
