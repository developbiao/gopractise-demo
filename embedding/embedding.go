package main

import "fmt"

// base
type base struct {
	num int
}

// describe
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container
type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

}
