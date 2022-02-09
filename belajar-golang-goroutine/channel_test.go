package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "anshoryihsan"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "ihsan"
		fmt.Println("selesai mengirim data ke channel")
	}()
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

////channe in & out
func OnlyIn(channel chan<- string) { //in
	time.Sleep(2 * time.Second)
	channel <- "ihsan"
}

func OnlyOut(channel <-chan string) { //out
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

////end channe in & out

////range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			channel <- "pengulanan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("manerima data", data)
	}
	fmt.Println("selesai")
}

////range channel

////select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string, 0)
	channel2 := make(chan string, 0)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

////end select channel
func TestSelectDefaultChannel(t *testing.T) {
	channel1 := make(chan string, 0)
	channel2 := make(chan string, 0)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("loading...")
		}
		if counter == 2 {
			break
		}
	}
}

////default select channel //karena ketika melakukan select shannel tidak mendapatkan nilainya atau kosong
////maka akan deadlock atau error, maka dari itu untuk mengakalinya menggunakan default select

////end default select channel
