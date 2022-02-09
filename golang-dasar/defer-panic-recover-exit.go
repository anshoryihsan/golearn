package main

import "fmt"

//defer
func logging() {
	fmt.Println("selesai memanggil function")
}
func runApplication(val int) {
	defer logging()
	fmt.Println("run application")
	result := 10 / val
	fmt.Println("result", result)
}

//end defer
//panic
func endApp() {
	fmt.Println("Aplikasi selesai")
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("aplikasi error")
	}
	fmt.Println("aplikasi berjalan")
}

//end panic
//recover
func endApp2() {
	message := recover()
	if message != nil {
		fmt.Println("error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}
func runApp2(error bool) {
	defer endApp2()
	if error {
		panic("aplikasi errore")
	}
	fmt.Println("aplikasi berjalan")
}

//end recover

func main() {
	// runApplication(0)
	// runApp(true)
	runApp2(true)
	fmt.Println("ihsan")
}
