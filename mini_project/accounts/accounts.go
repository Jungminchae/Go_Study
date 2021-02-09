package accounts

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
func (a Account) Deposit(amount int) {
	a.balance += amount
}

// Balance method??
func (a Account) Balance() int {
	return &a.balance
}
