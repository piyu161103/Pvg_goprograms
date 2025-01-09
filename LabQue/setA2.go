package main

import "fmt"

func main() {
    var number int
    fmt.Print("Enter a number: ")
    fmt.Scan(&number)

    if number%2 == 0 {
        fmt.Println(number, "is Even")
    } else {
        fmt.Println(number, "is Odd")
    }
}