package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Hello,Go,World"
	fmt.Print(strings.Split(sampleString, ",")) // Output: ["Hello" "Go" "World"]
}