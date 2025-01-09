package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter the number of terms in the Fibonacci series: ")
    fmt.Scan(&n)

    if n <= 0 {
        fmt.Println("Please enter a positive integer.")
        return
    }

    a, b := 0, 1
    fmt.Print("Fibonacci series: ")
    for i := 1; i <= n; i++ {
        fmt.Printf("%d ", a)
        a, b = b, a+b
    }
    fmt.Println()
}