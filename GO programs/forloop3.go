package main

import "fmt"

func main() {
    var n, a, b int
    fmt.Print("Enter the number of terms: ")
    fmt.Scan(&n)

    a, b = 0, 1
    fmt.Print("Fibonacci Series: ")
    for i := 0; i < n; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
    fmt.Println()
}