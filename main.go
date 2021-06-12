package main

import (
	"fmt"
	"strings"
)

// naked return
// 리턴 타입에 변수를 바로 선언 (생성)
// 함수가 끝나면 반환
func lenAndUpper(name string) (length int, uppercase string) {
	// defer 함수를 끝나고 실행
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}


func main() {
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
}
