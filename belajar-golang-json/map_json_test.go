package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"apple mac book","ImageURL":"http://blabla.com/image.png","price":400}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "apple mac book",
		"price": 2000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
