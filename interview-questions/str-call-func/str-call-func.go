package main

import "fmt"
import "reflect"

type Animal struct {
} 

func (a *Animal) Say() {
    fmt.Println("What fox say? ding ding ding~")
}

func main() {
    a := Animal{}
    reflect.ValueOf(&a).MethodByName("Say").Call([]reflect.Value{})
    fmt.Println("Main func done!")
}

