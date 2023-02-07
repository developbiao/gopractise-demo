package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

// Implemented using value receiver
func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

// Implemented using pointer receiver
func (a *Address) Describe() {
	fmt.Printf("State %s Country %s \n", a.state, a.country)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	case string:
		fmt.Println("I am a string type")
	default:
		fmt.Printf("Unkonw type\n")
	}
}

func main() {
	var d1 Describer
	p1 := Person{"Sam", 22}
	d1 = p1
	d1.Describe()

	p2 := Person{"James", 35}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	/*
	 compilation error if the floowing line uncommend
	 Describe method has pointer receiver
	 d2 = a
	*/
	d2 = &a
	d2.Describe()

	findType(d2)
}
