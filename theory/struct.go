package main

import "fmt"

//go 에서는 struct가 매우 중요
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"ramen", "kimchi"}
	minchae := person{name: "minchae", age: 28, favFood: favFood}
	fmt.Println(minchae)
}
