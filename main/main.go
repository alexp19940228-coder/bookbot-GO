package main

import (
	"fmt"
	"os"

	"github.com/alexp19940228-coder/stats"
)

func main() {
	book, err := os.ReadFile("../books/frankenstein.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(book)
	//	fmt.Println(text)
	fmt.Println(stats.WordCount(text))
	countWord := stats.CountChar(text)

	for ch, count := range countWord {
		fmt.Printf("%q : %d\n", ch, count)
	}
}
