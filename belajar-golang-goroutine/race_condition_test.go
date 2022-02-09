package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//race condition adalah kondisi diamana ketika banyak go routine yang meng akses variabel {{variabel x}} berebutan/mendahului
func TestRaceCondition(t *testing.T) {
	//solusinya tambahkan mutex agar kondisi variabel yang di akses go routine lebih aman dilakukan locking dan unlocking dengan mutex
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
				// x = 200 + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
