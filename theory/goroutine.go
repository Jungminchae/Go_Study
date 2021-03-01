package main

import (
	"fmt"
	"time"
)

// goroutine은 메인 함수가 실행되고 있을 때만 작동 됨
func main() {
	go sexyCount("minchae")
	sexyCount("fuck!")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is_sexy", i)
		time.Sleep(time.Second)
	}
}
