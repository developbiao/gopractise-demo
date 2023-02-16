package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			fmt.Println("finished reading file")
			break
		}
		if err != nil {
			fmt.Printf("Error %s reading file", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}

	fmt.Println("Main func done!")
}
