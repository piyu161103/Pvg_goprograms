package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Go Go Go!"
	fmt.Println(strings.Count(sampleString, "Go")) // Output: 3
}