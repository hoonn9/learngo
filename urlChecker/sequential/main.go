package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	// panic(Error) 초기화 되지 않은 map에 입력할 경우

	// var results = map[string]string{}
	var results = make(map[string]string)

	
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.naver.com/",
		"https://www.amazon.com/",
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		}


	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range(results) {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	println("checking... ",url)
	res, err := http.Get(url)

	// 400 부터 problem
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return  nil
}