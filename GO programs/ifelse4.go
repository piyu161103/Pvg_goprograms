package main

import (
    "fmt"
    "unicode"
)

func main() {
    var ch rune
    fmt.Print("Enter a character: ")
    fmt.Scanf("%c", &ch)

    if unicode.IsUpper(ch) {
        fmt.Println("The character is uppercase.")
    } else {
        fmt.Println("The character is lowercase.")
    }
}