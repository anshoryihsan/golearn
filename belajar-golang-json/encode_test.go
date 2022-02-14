package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "ihsan",
		LastName:   "anshory",
		MiddleName: "muhammad",
	}
	encoder.Encode(customer)
}
