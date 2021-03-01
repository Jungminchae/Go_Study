package main

import (
	"fmt"
	"time"
)

// goroutine은 메인 함수가 실행되고 있을 때만 작동 됨
func main() {
	c := make(chan bool)
	people := [5]string{"minchae", "minchae2", "minchae3", "minchae4", "minchae5"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- person + "is sexy"
}
