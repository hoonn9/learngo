package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// if  age < 18 {
	// 	return false
	// } else {
	// 	return true
	// }

	// if 안에 변수 선언 가능
	// 조건문에서만 사용하는 변수라는 것을 표현하기에 좋음
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}


func main() {
	fmt.Println(canIDrink(16))
}
