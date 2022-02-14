package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}
type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSSINObject(t *testing.T) {
	customer := Customer{
		FirstName:  "ihsan",
		MiddleName: "anshory",
		LastName:   "muhammad",
		Age:        27,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
