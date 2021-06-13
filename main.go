package main

import (
	"fmt"
	"log"

	"github.com/hoonn9/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("hoon")

	// 자동으로 복사본을 생성해서 poniter가 달라짐 => 타입에 pointer로 해결
	account.Deposit(10)
	fmt.Println(account.Balance())

	// Error handling
	// go에는 exception이 없고 강제로 에러를 체크하도록 함
	err := account.Withdraw(20)
	if err != nil {
		// print 후 프로그램 종료
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
