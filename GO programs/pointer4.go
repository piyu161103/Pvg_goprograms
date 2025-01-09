package main

import "fmt"

func multiply(a, b *int) int {
    return *a * *b
}

func main() {
    var x, y int = 4, 5
    result := multiply(&x, &y)

    fmt.Println("Result of multiplication:", result)
}