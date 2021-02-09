package main

import "fmt"

// for문은 for밖에 없음

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func superAdd2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func superAdd3(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)

	superAdd2(1, 2, 3, 4, 5, 6)

	result := superAdd3(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
