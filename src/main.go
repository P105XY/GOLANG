package main

import (
	"fmt"

	"./MyDict"
)

func main() {
	mdic := MyDict.Dictionary{}
	mdic["Hello"] = "HELLLO"
	fmt.Println(mdic)
}
