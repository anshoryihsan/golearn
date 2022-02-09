package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{"febuary": 31, "mei": 30}
	chicken["januari"] = 31
	fmt.Print(chicken)
	//for range
	for key, val := range chicken {
		fmt.Println(key, " \t", val)
	}
	//find item using key
	var value, isExist = chicken["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
	//combine slice and map
	var chickens2 = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		{"name": "chicken grey", "gender": "shemale"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken2 := range chickens2 {
		fmt.Println(chicken2["gender"], chicken2["name"])
	}
}
