package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d
	// kembalikan 2 nilai
	return area, circumference
}

func getFullName() (string, string, string) {
	return "muhammad", "ihsan", "anshory"
}
func main() {
	firstName, midName, lastName := getFullName()
	fmt.Println(firstName, midName, lastName)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}
