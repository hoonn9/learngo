package accounts

import "errors"

// 속성도 public 하려면 대문자

// Account struct
// type Account struct {
// 	Owner 	string
// 	Balance int
// }

// acount struct
type Account struct {
	owner 	string
	balance int
}

// go의 contructor
// NewAccount creates Account
func NewAccount(owner string) *Account {
	// 새로운 객체 만들고 메모리 주소를 반환
	account := Account{owner: owner, balance: 0}	
	return &account
}

var errNoMoney = errors.New("Can't withdraw")


// Method
// recevier
// 타입 이름의 앞글자를 따서 소문자로 시작하는 이름으로 짓는다. (Account -> a...)
// func (a Account) Deposit(amount int) {
// 	a.balance += amount
// }

// 객체 접근할 때 자동 복사본을 사용하지 않으려면 타입에 pointer
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of account
func (a Account) Balance() int {
	return a.balance
}

// Error
// error는 error | nil (null)
// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	
	a.balance -= amount
	return nil
}