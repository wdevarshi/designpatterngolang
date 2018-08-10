package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	author
	title   string
	content string
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	//fmt.Println("Author: ", p.author.fullName())
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.author.bio)
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		author1,
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
	}
	post1.details()
}
