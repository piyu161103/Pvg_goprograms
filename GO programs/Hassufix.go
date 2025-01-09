package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Hello, Go World!"
	fmt.Println(strings.HasSuffix(sampleString, "World!")) // Output: true
}