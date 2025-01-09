package main

import "fmt"

func main() {
    var n, sum int
    fmt.Print("Enter the number of elements: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    fmt.Println("Enter the elements:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
        sum += arr[i]
    }

    average := float64(sum) / float64(n)
    fmt.Println("Sum of elements:", sum)
    fmt.Println("Average of elements:", average)
}