package accounts

import (
	"errors"
	"fmt"
)

//struct를 만들거임

// Account struct
type Account struct {
	// 소문자 시작이면 private이라 접근할 수 없음
	owner   string
	balance int
}

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance method??
func (a Account) Balance() int {
	return a.balance
}

var errNomoney = errors.New("Can't withdraw. you are poor!")

// Withdraw money
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNomoney
	}
	a.balance -= amount
	// nil == none 같은거
	return nil
}

// Change owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// python에 있는  __str__같은거
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account \nHas: ", a.Balance())
}
