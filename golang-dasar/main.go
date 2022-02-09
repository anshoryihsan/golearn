package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//struct function
func (customer Customer) sayHi(name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}
func (val Customer) huu(name string) {
	fmt.Println("huu from", val.Name)
}

//end struct function

func main() {
	var cust Customer
	cust.Name = "ihsan"
	cust.Address = "bandung"
	cust.Age = 27
	fmt.Println(cust)
	//struct literal
	cust2 := Customer{
		Name:    "joko",
		Address: "jakarta",
		Age:     30,
	}
	fmt.Println(cust2)
	cust3 := Customer{"budi", "purwakarta", 30}
	fmt.Println(cust3)

	// struct function
	cust.sayHi("anshory")
	cust.huu("anshory")
	// end struct function
}
