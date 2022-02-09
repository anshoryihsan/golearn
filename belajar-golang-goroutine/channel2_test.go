package belajargolanggoroutine

import (
	"fmt"
	"testing"
)

// func sayHelloTo(who string) {
// 	// var message = make(chan string)
// 	var data = fmt.Sprint("helo %s", who)
// 	message <- data
// }

func TestChan(t *testing.T) {
	var message = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("helo %s", who)
		message <- data
	}
	go sayHelloTo("ihsan")
	go sayHelloTo("anshory")
	go sayHelloTo("muhammad")

	var message1 = <-message
	fmt.Println(message1)
	var message2 = <-message
	fmt.Println(message2)
	var message3 = <-message
	fmt.Println(message3)
}
