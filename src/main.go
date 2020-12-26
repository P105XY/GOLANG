package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	//REMEMBER : map 생성 시, 선언과 동시에 초기화 필수.
	//어떤 값이든 초기화를 시켜야, 나중에 따로 대입이나 재정의가 가능하다.
	var results = make(map[string]string)

	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err)
		return errRequestFailed
	}
	return nil
}
