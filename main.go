package main

import (
	"fmt"
)


func main() {
	// limit item
	names := [5]string{"hoon", "kozel", "blang"}
	names[3] = "ddd"
	names[4] = "agdsg"

	// slice (not length)
	names2 := []string{"hoon", "kozel", "blang"}
	// names2[3] = "ddd" // Error ❌

	// 기존 array 을 변경하지 않고 새로 만들어서 반환
	names2 = append(names2, "good")

	fmt.Println(names2)
}
