//menggabungkan interface
package main

import (
	"log"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}
func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}
func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunRuang hitung = &kubus{4}

	log.Println("volume kubus : ", bangunRuang.volume())
	log.Println("keliling kubus : ", bangunRuang.keliling())
	log.Println("luas kubus : ", bangunRuang.luas())
}
