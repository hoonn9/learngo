package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)

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
		go hitURL(url, c)
	}
	
	for i:=0; i<len(urls); i++ {
		result := <- c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// chan<-   => Send only


func hitURL(url string, c chan<- requestResult) {
	res, err := http.Get(url)
	status := "OK"
	
	if err != nil || res.StatusCode >= 400 {
		c <- requestResult{url: url, status: "FAILED"}
	}
	c <- requestResult{url: url, status: status}
	
}