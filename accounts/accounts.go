package accounts

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

