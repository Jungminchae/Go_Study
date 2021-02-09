package main

import "fmt"

func main() {
	// python의 dict같은 느낌
	// variable := map[key_type]value_type{key:value ...}
	minchae := map[string]string{
		"name": "minchae", "age": "15"}
	fmt.Println(minchae)

	for key, value := range minchae {
		fmt.Println(key, value)
	}
}
