package main

import (
	"fmt"
)


func main() {
	// Map  (js object 와 비슷하지만 또 다름)
	//map[key][value]{입력}
	hoon := map[string]string{
		"name": "hoon",
		"age": "12",
	}

	fmt.Println(hoon)

	for key, value := range hoon {
		fmt.Println(key, value)
	}
}
