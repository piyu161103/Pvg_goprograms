package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)

    if a > b {
        fmt.Println("The first number is greater.")
    } else {
        fmt.Println("The second number is greater.")
    }
}