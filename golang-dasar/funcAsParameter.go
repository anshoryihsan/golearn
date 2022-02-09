package main

import "fmt"

type FIlter func(string) string

func sayHelloFilter(name string, filter func(string) string) { //func type declaration
	// func sayHelloFilter(name string, filter func(string) string) { //func normal
	nameFilered := filter(name)
	fmt.Println("hello", nameFilered)
}
func filtered(name string) string {
	if name == "anjing" {
		return "..."
	}
	return name
}

func main() {
	sayHelloFilter("ihsan", filtered)
	// say := sayHelloFilter
	// say("ihsan", filter())
}
