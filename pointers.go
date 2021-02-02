package main

import "fmt"

func main() {
	// 메모리 주소 보는법
	a := 2
	b := 5
	c := &a
	fmt.Println(&a, &b)
	// see through
	// 메모리주소 값에 들어있는 값을 보는 것
	fmt.Println(*c)
	// c는 a의 주소값을 받고있기 때문에 a의 값이 바뀌면 *c 도 바뀜
	a = 20
	fmt.Println(*c)
	// 주소값을 가지고 원래 값을 변경시킬수도 있음
	*c = 30
	fmt.Println(a)
}
