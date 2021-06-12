package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	// Struct  (js object 와 비슷하고 Map 보다 유연)
	// Go 는 constructor 가 없음. 직접 실행해야 함

	// 구조체 선언법
	// 1.
	favFood := []string{"kimchi", "coffee"}
	// hoon := person{"hoon", 26, favFood}

	// 2.
	hoon := person{name: "hoon", age: 26, favFood: favFood}
	fmt.Println(hoon.name)
}
