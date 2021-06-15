package main

import (
	"fmt"
	"time"
)

func main() {
	// 함수 앞에 go만 써주면 병행성 실행

	go sexyCount("hoon")
	sexyCount("kozel")

	// go routines 은 메인 함수가 끝나면 종료 됨 (기다려 주지 않음)

	// go sexyCount("kozel")
	// go sexyCount("hoon")
}

func sexyCount(person string) {
	// top down
	for i:=0; i<10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}