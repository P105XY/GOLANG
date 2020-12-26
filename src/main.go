package main

import (
	"fmt"

	"./MyDict"
)

func main() {
	mdic := MyDict.Dictionary{}

	baseWord := "Hello"
	mdic.Add(baseWord, "first")
	mdic.Search(baseWord)
	mdic.Delete(baseWord)
	word, err := mdic.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(word)
}
