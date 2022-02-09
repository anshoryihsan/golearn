package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayhi() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i]
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayhi()
	var name = s1.getNameAt(2)
	fmt.Println(name)
}
