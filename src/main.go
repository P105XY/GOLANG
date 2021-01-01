package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	//REMEMBER : map 생성 시, 선언과 동시에 초기화 필수.
	//어떤 값이든 초기화를 시켜야, 나중에 따로 대입이나 재정의가 가능하다.
	var results = make(map[string]string)
	defer fmt.Println("func end")

	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
	}
	channel := make(chan string)
	people := []string{"nico", "flynn", "adlf", "guuy"}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	for _, person := range people {
		go isTrue(person, channel)
	}
	for url, result := range results {
		fmt.Println(url, result)
	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
}

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err)
		return errRequestFailed
func Count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, " counted", i)
		time.Sleep(time.Second)
	}
	return nil
}

func isTrue(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " it true"
}