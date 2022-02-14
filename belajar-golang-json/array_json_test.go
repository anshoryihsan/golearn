package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "ihsan",
		MiddleName: "anshory",
		LastName:   "muhammad",
		Hobbies:    []string{"gaming", "reading", "coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"ihsan","MiddleName":"anshory","LastName":"muhammad","Age":0,"Married":false,"Hobbies":["gaming","reading","coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "ihsan",
		Addresses: []Address{
			{
				Street:     "jalan belum ada",
				Country:    "indonesia",
				PostalCode: "99",
			},
			{
				Street:     "jalan belum ada",
				Country:    "indonesia",
				PostalCode: "88",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"ihsan","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"jalan belum ada","Country":"indonesia","PostalCode":"99"},{"Street":"jalan belum ada","Country":"indonesia","PostalCode":"88"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"jalan belum ada","Country":"indonesia","PostalCode":"99"},{"Street":"jalan belum ada","Country":"indonesia","PostalCode":"88"}]`
	jsonBytes := []byte(jsonString)

	adresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, adresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(adresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "jalan belum ada",
			Country:    "indonesia",
			PostalCode: "99",
		},
		{
			Street:     "jalan belum ada",
			Country:    "indonesia",
			PostalCode: "88",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
