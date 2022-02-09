package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoo(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("muhammad")
	pool.Put("ihsan")
	pool.Put("anshory")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("selesai")
}
