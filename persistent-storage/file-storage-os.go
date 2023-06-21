package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("你好中国🇨🇳!\n")
	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file \n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)

	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(read2, string(read2))
}
