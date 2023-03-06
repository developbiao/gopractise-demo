package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "Save the people"
	}
	return
}

func main() {
	var people People = &Student{}
	fmt.Println(people.Speak("Hello"))

	fmt.Println("Main func done!")
}
