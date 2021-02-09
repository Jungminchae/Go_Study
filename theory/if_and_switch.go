package main

import "fmt"

func canIDrink(age int) bool {
	// variable expression
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	// 값을 체크해주는 문법
	// switch도 variable expression 가능
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))

	fmt.Println(canIDrink2(25))
}
