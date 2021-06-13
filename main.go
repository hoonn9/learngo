package main

import (
	"fmt"

	"github.com/hoonn9/learngo/accounts"
)

func main() {
	// 생성자 없이 만드는 법
	// account := banking.Account{Owner: "hoon", Balance: 1000000}
	// account.Owner = "pepe";
	// fmt.Println(account)

	// private 여도 constructor 함수 선언해서 객체 생성 가능
	accounts := accounts.NewAccount("hoon")
	// accounts.balance = 12 // Error private라 변경 불가
	fmt.Println(accounts)
}
