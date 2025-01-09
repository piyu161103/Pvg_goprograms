package main

import "fmt"

func main() {
    var num, reversed int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    for num > 0 {
        digit := num % 10
        reversed = reversed*10 + digit
        num /= 10
    }
    fmt.Println("Reversed number is", reversed)
}