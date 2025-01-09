package main

import "fmt"

func main() {
    var ch string
    fmt.Print("Enter a character: ")
    fmt.Scan(&ch)

    if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" {
        fmt.Println("It's a vowel.")
    }
}