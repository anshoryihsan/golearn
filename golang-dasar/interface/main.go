package main

import (
	"log"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (c lingkaran) jariJari() float64 {
	return c.diameter / 2
}
func (c lingkaran) luas() float64 {
	return math.Pi * math.Pow(c.jariJari(), 2)
}
func (c lingkaran) keliling() float64 {
	return math.Pi * c.jariJari()
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}
func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var bangunDatar hitung
	bangunDatar = persegi{10.0}
	log.Println("luas persegi : ", bangunDatar.luas())
	log.Println("luas persegi : ", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	log.Println("luas lingkaran : ", bangunDatar.luas())
	log.Println("luas lingkaran : ", bangunDatar.keliling())
}
