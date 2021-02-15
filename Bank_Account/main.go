package main

import (
	"fmt"

	"accounts"
)

func main() {
	account := accounts.NewAccount("minchae")
	account.Deposit(10)
	// error가 있으면 직접 낼 수 있도록 코딩을 해줘야함
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}
