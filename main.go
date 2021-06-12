package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}

	// 마치 if문처럼 사용가능
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}

	// 마치 if문처럼 사용가능
	switch koreanAge := age + 2; koreanAge {
		case  10:
			return false
		case  18:
			return true
	}
	return false
}


func main() {
	fmt.Println(canIDrink(16))
}
