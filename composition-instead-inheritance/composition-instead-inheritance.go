package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) detail() {
	fmt.Println("Title:", b.title)
	fmt.Println("Content:", b.content)
	fmt.Println("Bio:", b.bio)
	fmt.Println("Author:", b.author.fullName())
}

type website struct {
	blogPosts []blogPost
}

func (w website) contents() {
	for _, v := range w.blogPosts {
		v.detail()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}

	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to struct",
		author1,
	}

	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a paraller one",
		author1,
	}

	w := website{[]blogPost{blogPost1, blogPost2, blogPost3}}
	w.contents()
}
