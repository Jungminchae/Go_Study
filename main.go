package main

// 컴파일을 위해서 필요함 main
// 컴파일이 필요없다면 노필요

// 임포트는 이렇게
import (
	"fmt"
	"something"
	"strings"
)

// 함수만들기
func multiply(a, b int) int {
	return a * b
}

// 멀티 리턴
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 다중 인자 받는법
func repeatMe(words ...string) {
	fmt.Println(words)
}

//nake return
// 리턴을 미리 정해주는 것
func lenAndUpper2(name string) (length int, uppercase string) {
	// 새로운 변수에 할당하는 것이 아닌 변수를 업데이트 하는 것
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// defer 함수를 실행하고 실행되는 것
func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// 메인 함수 있어야함
// 끝에 ; 필요없음
func main() {
	fmt.Println("Hello World")
	// 변수 : 타입을 알려줘야함
	// JS와 같이 상수인 변수는 const
	const name string = "minchae"
	fmt.Println(name)
	// var variable string = "VAR"  이런 축약형은 함수 안에서만 사용가능
	// variable := "VAR"  // := 이놈이 변수 타입을 찾아줄 것이야

	//multiply 함수 사용
	fmt.Println(multiply(2, 2))

	//멀티리턴
	// 하나만 받을라하면 에러남 , python에서 쓰던것 처럼 _ 사용하자, 안받을거면
	totalLength, upperName := lenAndUpper("minchae")
	fmt.Println(totalLength, upperName)

	//다중인자 받는 함수
	repeatMe("와", "이거", "개지리는", "언어네?")

	// naked return
	fmt.Println(lenAndUpper2("minchaes"))

	// defer
	fmt.Println(lenAndUpper3("jung"))

	something.SayHello()
}
