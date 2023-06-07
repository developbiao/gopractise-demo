package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func main() {
	var i interface{} = new(Student)
	// var i interface{} = (*Student)(nil)
	// var i interface{}
	// var i interface{} = &Student{}
	// Safe assert
	// s, ok := i.(Student)
	// if ok {
	// 	fmt.Println(s)
	// }
	fmt.Printf("%p %v\n", &i, i)
	judge(i)
}

func judge(v interface{}) {
	fmt.Printf("%p %v\n", &v, v)

	switch v := v.(type) {
	case nil:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("nil type[%T] %v\n", v, v)
	case Student:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("Studnt type[%T] %v\n", v, v)
	case *Student:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("*Student type[%T] %v\n", v, v)
	default:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("unkonw\n")
	}
}