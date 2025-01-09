package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Learn", "Go", "Easily"}
	fmt.Println(strings.Join(words, " ")) // Output: "Learn Go Easily"
}