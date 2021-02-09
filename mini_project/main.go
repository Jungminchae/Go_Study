package main

import (
	"fmt"

	"accounts"
)

func main() {
	account := accounts.NewAccount("minchae")
	fmt.Println(account)
}
