package main

import "fmt"

//factorial normal
func factorialLoop(val int) int {
	result := 1
	for i := val; i > 0; i-- {
		result *= i
	}
	return result
}

//factorial recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	fmt.Println(factorialLoop(5))
	result := factorialRecursive
	fmt.Println(result(5))
}
