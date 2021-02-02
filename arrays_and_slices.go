package main

import "fmt"

func main() {
	// go 에서 array를 표현하려면 [크기]타입{데이터}
	// go 에서 제한없는 array는 slice, 크기 부분을 비워놓으면 됨
	names := [5]string{"Jung", "Min", "Chae"}
	fmt.Println(names)
	ages := []int{15, 25, 35}
	ages = append(ages, 35)
	fmt.Println(ages)
}
