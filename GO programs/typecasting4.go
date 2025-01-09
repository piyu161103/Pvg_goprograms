package main

import "fmt"

func main() {
	var char rune='p'
	var str string=string(char)
	fmt.Println("Rune value:",char)
    fmt.Println("converted string value:",str)
	var newRune rune=rune(str[0])
	fmt.Println("converted rune value:",newRune)
}
