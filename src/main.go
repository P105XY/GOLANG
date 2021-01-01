package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("func end")

	channel := make(chan string)
	people := []string{"nico", "flynn", "adlf", "guuy"}

	for _, person := range people {
		go isTrue(person, channel)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}

}

func Count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, " counted", i)
		time.Sleep(time.Second)
	}
}

func isTrue(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " it true"
}

//comp url
