package main

import "fmt"

func main() {
    var base, exp, result int
    fmt.Print("Enter base and exponent: ")
    fmt.Scan(&base, &exp)

    result = 1
    i := 0
    for {
        if i >= exp {
            break
        }
        result *= base
        i++
    }
    fmt.Println("Result is", result)
}