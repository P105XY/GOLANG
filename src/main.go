package main

import (
	"fmt"

	"./MyDict"
)

func main() {
	mdic := MyDict.Dictionary{"first": "firstWord"}
	mditn, err := mdic.Search("Second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mditn)
	}
}
