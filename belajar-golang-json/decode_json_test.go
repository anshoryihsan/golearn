package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customers struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"ihsan","MiddleName":"anshory","LastName":"muhammad","Age":27,"Married":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customers{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
}
