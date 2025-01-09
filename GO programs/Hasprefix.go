package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Hello, Go World!"
	fmt.Println(strings.HasPrefix(sampleString, "Hello")) // Output: true
}