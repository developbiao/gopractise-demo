package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protobuf_using_go/study"
)

func main() {
	myBook := &study.Book{
		Name: "Animal Farm",
		Isbn: 103,
		Author: &study.Author{
			Name:             "George Orwell",
			YearOfPublishing: 1935,
		},
	}

	data, err := proto.Marshal(myBook)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)

	myNewBook := &study.Book{}
	err = proto.Unmarshal(data, myNewBook)
	if err != nil {
		log.Fatal("Unmarshaling error:", err)
	}
	fmt.Println(myNewBook.GetName())
	fmt.Println(myNewBook.GetIsbn())
	fmt.Println(myNewBook.Author.GetName())
	fmt.Println(myNewBook.Author.GetYearOfPublishing())
}
