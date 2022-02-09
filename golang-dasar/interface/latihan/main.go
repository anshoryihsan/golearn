package main

import (
	"fmt"
)

/*
===membuat patern star menggunakan interface
*
**
***
****
*****
			....*
			...**
			..***
			.****
			*****
===membuat patern star menggunakan interface
**/

type paternStar interface {
	normal() string
	upNormal() string
}
type stars struct {
	str int
}

func (s stars) normal() string {
	result := ""
	for i := 0; i < s.str; i++ {
		for j := 0; j <= i; j++ {
			result += "*"
		}
		result += "\n"
	}
	return result
}
func (s stars) upNormal() string {
	result := ""
	for i := 0; i < s.str; i++ {
		for j := s.str - 1; j >= i; j-- {
			result += " "
		}
		for j := 0; j <= i; j++ {
			result += "*"
		}
		result += "\n"
	}
	return result
}
func main() {
	var s paternStar = stars{5}
	fmt.Print(s.normal())
	fmt.Println()
	fmt.Print(s.upNormal())
	// log.Print(starNormal.normal())
}
