package something

import (
	"fmt"
)

// 대문자로 시작하면 다른 곳으로  export 될 수 있음
func SayHello() {
	fmt.Println("Hello")
}

func sayBye() {
	fmt.Println("Bye")
}
