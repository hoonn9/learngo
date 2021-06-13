package main

import (
	"fmt"

	"github.com/hoonn9/learngo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	
}
