package main

import (
	"fmt"
	"time"
)

// Channel 은 Go routines 과 pipe와 같은 형태로 데이터를 주고 받는다.
func main() {

	// Channel 선언
	// chan (주고 받은 정보 타입)
	channel := make(chan string)
	people := [5]string{"hoon", "kozel", "react", "typescript", "linux"}

	for _, person := range people {
		go isSexy(person, channel)
	}
	// channel 에서 데이터 받기
	// 받을 때 까지 기다림
	// result := <- channel

	// block operation 함수 멈춤 (<-)
	// fmt.Println("Waiting messages...")
	// fmt.Println("Received this message: ", <- channel)
	// fmt.Println("Received this message: ", <- channel)

	for i:=0; i<len(people); i++ {
		fmt.Println(<- channel)
	}
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 10)
	// channel 에 데이터 보내기
	channel <- person + " is sexy"
}