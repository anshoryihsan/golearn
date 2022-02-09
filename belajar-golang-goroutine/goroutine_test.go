package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func DispayNumber(number int) {
	fmt.Println("Display", number)
}
func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		go DispayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
