package main

import (
	"fmt"

	"github.com/hoonn9/learngo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}

	baseWord := "hello"
	definition := "Greeting"

	err := dictionary.Add(baseWord, definition)
	if err != nil {
		fmt.Println(err)
	}
	
	hello, _ := dictionary.Search("hello")
	fmt.Println(hello)

	dictionary.Update(baseWord, "Second")
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	baseWord2 := "hello2"

	dictionary.Add(baseWord2, "world")

	dictionary.Delete(baseWord2)
	word2, _ := dictionary.Search(baseWord2)
	fmt.Println(word2)
}
