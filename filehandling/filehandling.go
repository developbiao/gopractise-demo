package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed test.txt
var contents []byte

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("Value of fpath is", *fptr)

	if info, err := os.Stat(*fptr); err != nil {
		fmt.Println("File does'not exists")
		return

	} else {
		fmt.Println("File information:", info)
	}

	contents, err := os.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reding error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))

	fmt.Println("Contents with embed:", string(contents))

	fmt.Println("Main func done!")
}
